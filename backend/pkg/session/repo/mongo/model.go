package mongo

import (
	"time"

	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type session struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Start    time.Time          `bson:"start"`
	End      time.Time          `bson:"end"`
}

func (s *session) toServiceSession() *model.Session {
	return &model.Session{
		Name:  s.Name,
		Start: s.Start,
		End:   s.End,
		ID:    s.ObjectID.Hex(),
	}
}

func fromServiceSession(s *model.Session) (*session, error) {
	sess := &session{
		Name:  s.Name,
		Start: s.Start,
		End:   s.End,
	}
	if s.ID != "" {
		id, err := primitive.ObjectIDFromHex(s.ID)
		if err != nil {
			return nil, err
		}
		sess.ObjectID = id
	}
	return sess, nil
}
