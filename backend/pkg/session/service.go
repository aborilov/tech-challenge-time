package session

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/model"
)

type service struct {
	repo model.Repository
}

// NewService create new service that implements Service interface
func NewService(repo model.Repository) model.Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context, filter model.Filter) ([]*model.Session, error) {
	return s.repo.List(ctx, filter)
}

func (s *service) Get(ctx context.Context, id string) (*model.Session, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) StartNewSession(ctx context.Context) (*model.Session, error) {
	session, err := s.GetCurrentSession(ctx)
	if err == nil {
		return nil, model.ErrSessionInProgress{ID: session.ID}
	}
	fmt.Println(session)
	if errors.As(err, &model.ErrNoSessionInProgress{}) {
		session := model.Session{Start: time.Now()}
		return s.repo.Add(ctx, &session)
	}
	return nil, err
}
func (s *service) StopCurrentSession(ctx context.Context) (*model.Session, error) {
	session, err := s.GetCurrentSession(ctx)
	if err != nil {
		return nil, err
	}
	session.End = time.Now()
	return s.repo.Update(ctx, session)
}
func (s *service) GetCurrentSession(ctx context.Context) (*model.Session, error) {
	f := model.Filter{WithoutEnd: true}
	sessions, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}
	if len(sessions) == 0 {
		return nil, model.ErrNoSessionInProgress{}
	}
	// should be only one running session
	session := sessions[0]
	return session, nil

}
func (s *service) Update(ctx context.Context, session *model.Session) (*model.Session, error) {
	return s.repo.Update(ctx, session)
}
