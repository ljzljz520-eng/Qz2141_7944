package query

import (
	"sort"
	"strings"
	"training29/model"
	"training29/storage"
)

type Query struct{ Store *storage.Store }

func New(s *storage.Store) *Query             { return &Query{Store: s} }
func (q *Query) All() ([]model.Record, error) { return q.Store.ListRecords() }
func (q *Query) ByStatus(status string) ([]model.Record, error) {
	rs, e := q.All()
	if e != nil {
		return nil, e
	}
	out := rs[:0]
	for _, r := range rs {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out, nil
}
func (q *Query) Search(term string) ([]model.Record, error) {
	rs, e := q.All()
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(term)
	out := []model.Record{}
	for _, r := range rs {
		if strings.Contains(strings.ToLower(r.ID), term) || strings.Contains(strings.ToLower(r.UserID), term) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (q *Query) Sorted(rs []model.Record) []model.Record {
	sort.Slice(rs, func(i, j int) bool { return rs[i].CreatedAt.Before(rs[j].CreatedAt) })
	return rs
}
