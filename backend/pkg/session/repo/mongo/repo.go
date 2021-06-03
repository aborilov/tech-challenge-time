package mongo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "tracker"
	collection = "sessions"
)

type repo struct {
	store      []*model.Session
	collection *mongo.Collection
}

// NewRepository creates mongo implementation of Repository interface
func NewRepository(client *mongo.Client) model.Repository {
	collection := client.Database(database).Collection(collection)
	return &repo{collection: collection, store: []*model.Session{}}
}

func (r *repo) Get(ctx context.Context, id string) (*model.Session, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	sess := session{}
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&sess)
	return sess.toServiceSession(), err
}

func (r *repo) List(ctx context.Context, filter model.Filter) ([]*model.Session, error) {
	f := bson.M{}
	if filter.WithoutEnd {
		f["end"] = time.Time{}
	}
	if !filter.StartAfter.IsZero() {
		f["start"] = bson.M{"$gte": primitive.NewDateTimeFromTime(filter.StartAfter)}
	}
	if !filter.EndBefore.IsZero() {
		f["end"] = bson.M{
			"$lte": primitive.NewDateTimeFromTime(filter.EndBefore),
			"$ne":  time.Time{},
		}
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"start": -1})

	cur, err := r.collection.Find(ctx, f, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	ss := []*model.Session{}
	for cur.Next(ctx) {
		s := session{}
		err := cur.Decode(&s)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s.toServiceSession())
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return ss, nil
}

func (r *repo) Add(ctx context.Context, session *model.Session) (*model.Session, error) {
	sess, err := fromServiceSession(session)
	if err != nil {
		return nil, err
	}
	sess.ObjectID = primitive.NewObjectID()
	res, err := r.collection.InsertOne(ctx, sess)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("can't get ID")
	}
	session.ID = oid.Hex()
	return session, nil
}

func (r *repo) Update(ctx context.Context, session *model.Session) (*model.Session, error) {
	sess, err := fromServiceSession(session)
	if err != nil {
		return nil, err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": sess.ObjectID}, sess)
	if err != nil {
		return nil, err
	}
	return session, nil
}
