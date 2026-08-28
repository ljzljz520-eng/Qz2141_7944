package query

import (
	"strings"
	"training29/model"
)

type Filter struct {
	Status, User, Event string
	FinalOnly           bool
}

func (f Filter) Match(r model.Record) bool {
	if f.Status != "" && r.Status != f.Status {
		return false
	}
	if f.User != "" && !strings.EqualFold(r.UserID, f.User) {
		return false
	}
	if f.Event != "" && r.EventID != f.Event {
		return false
	}
	if f.FinalOnly && !r.IsFinal() {
		return false
	}
	return true
}
func Apply(rs []model.Record, f Filter) []model.Record {
	out := make([]model.Record, 0, len(rs))
	for _, r := range rs {
		if f.Match(r) {
			out = append(out, r)
		}
	}
	return out
}
