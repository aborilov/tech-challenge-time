package model

import "fmt"

// ErrNoSessionInProgress raise when there is no tracking session in progress
type ErrNoSessionInProgress struct {
}

func (p ErrNoSessionInProgress) Error() string {
	return "there is no session in progress"
}

// ErrSessionInProgress raise when there is tracking session already in progress
type ErrSessionInProgress struct {
	ID string
}

func (p ErrSessionInProgress) Error() string {
	return fmt.Sprintf("there is already session in progress: %s", p.ID)
}
