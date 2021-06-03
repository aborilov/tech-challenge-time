package session

import (
	"time"

	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/model"
)

// swagger:response session
type session struct {
	// ID of the session
	ID string `json:"id"`
	// Name of the session
	Name string `json:"name"`
	// Start time of the session
	Start time.Time `json:"start"`
	// End time of the session
	End *time.Time `json:"end,omitempty"`
}

// swagger:response sessions
type sessions []*session

func (s *session) toSession() *model.Session {
	sess := &model.Session{
		ID:    s.ID,
		Name:  s.Name,
		Start: s.Start,
	}
	if s.End != nil {
		sess.End = *s.End
	}
	return sess
}

func fromSession(s *model.Session) *session {
	sess := &session{
		ID:    s.ID,
		Name:  s.Name,
		Start: s.Start,
	}
	if !s.End.IsZero() {
		sess.End = &s.End
	}
	return sess
}
