package model

import "time"

// Session - contains all tracking session fields
type Session struct {
	ID    string
	Name  string
	Start time.Time
	End   time.Time
}

// Filter - filter options that we can use with List method
type Filter struct {
	StartAfter time.Time
	EndBefore  time.Time
	WithoutEnd bool
}
