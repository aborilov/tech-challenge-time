package model

import "context"

// Service interface of session service
type Service interface {
	Get(context.Context, string) (*Session, error)
	List(context.Context, Filter) ([]*Session, error)
	StartNewSession(context.Context) (*Session, error)
	StopCurrentSession(context.Context) (*Session, error)
	GetCurrentSession(context.Context) (*Session, error)
	Update(context.Context, *Session) (*Session, error)
}

//go:generate mockery --name=Service --inpkg

// Repository interface of session repository
type Repository interface {
	Get(context.Context, string) (*Session, error)
	List(context.Context, Filter) ([]*Session, error)
	Add(context.Context, *Session) (*Session, error)
	Update(context.Context, *Session) (*Session, error)
}

//go:generate mockery --name=Repository --inpkg
