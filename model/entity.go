package model

import "time"

type Record struct {
	ID, UserID, EventID, Status, Notes string
	CreatedAt, UpdatedAt               time.Time
	Version                            int
}
type User struct {
	ID, Name, Email, Department string
	Active                      bool
	CreatedAt                   time.Time
}
type Event struct {
	ID, Code, Title, Location string
	Capacity                  int
	StartsAt                  time.Time
	Open                      bool
}
type Audit struct {
	ID, RecordID, Actor, Action, Detail string
	At                                  time.Time
}

func NewRecord(id, user, event string) Record {
	now := time.Now().UTC()
	return Record{ID: id, UserID: user, EventID: event, Status: "submitted", CreatedAt: now, UpdatedAt: now, Version: 1}
}
func (r Record) IsFinal() bool {
	return r.Status == "approved" || r.Status == "archived" || r.Status == "rejected"
}
func (r Record) CanTransition(next string) bool {
	switch r.Status {
	case "submitted":
		return next == "processing" || next == "rejected"
	case "processing":
		return next == "approved" || next == "rejected"
	case "approved":
		return next == "archived"
	}
	return false
}
func (u User) Valid() bool      { return u.ID != "" && u.Name != "" && u.Email != "" && u.Active }
func (e Event) Available() bool { return e.Open && e.Capacity > 0 && e.StartsAt.After(time.Now()) }
func (a Audit) Valid() bool     { return a.ID != "" && a.RecordID != "" && a.Action != "" }
