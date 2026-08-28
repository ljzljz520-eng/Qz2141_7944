package catalog

import (
	"sort"
	"strings"
	"training29/model"
)

type Roster struct {
	Event   model.Event
	Records []model.Record
}

func NewRoster(e model.Event) *Roster { return &Roster{Event: e, Records: []model.Record{}} }
func (r *Roster) Add(x model.Record) bool {
	if x.EventID != r.Event.ID || len(r.Records) >= r.Event.Capacity {
		return false
	}
	for _, old := range r.Records {
		if old.UserID == x.UserID {
			return false
		}
	}
	r.Records = append(r.Records, x)
	return true
}
func (r *Roster) Remove(id string) bool {
	for i, x := range r.Records {
		if x.ID == id {
			r.Records = append(r.Records[:i], r.Records[i+1:]...)
			return true
		}
	}
	return false
}
func (r *Roster) Find(term string) []model.Record {
	out := []model.Record{}
	for _, x := range r.Records {
		if strings.Contains(x.UserID, term) || strings.Contains(x.ID, term) {
			out = append(out, x)
		}
	}
	return out
}
func (r *Roster) Sorted() []model.Record {
	sort.Slice(r.Records, func(i, j int) bool { return r.Records[i].UserID < r.Records[j].UserID })
	return r.Records
}
