package reconcile

import (
	"fmt"
	"time"
	"training29/model"
	"training29/storage"
)

type Result struct{ Checked, Fixed, Conflicts int }

func Check(s *storage.Store) (Result, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return Result{}, e
	}
	r := Result{Checked: len(rs)}
	for _, x := range rs {
		if x.UpdatedAt.IsZero() {
			r.Conflicts++
		}
		if x.Status == "submitted" && time.Since(x.CreatedAt) > 72*time.Hour {
			r.Conflicts++
		}
	}
	return r, nil
}
func Repair(s *storage.Store, id string) error {
	return s.UpdateRecord(id, func(r *model.Record) error {
		if r.Status == "submitted" {
			r.Status = "processing"
			return nil
		}
		return fmt.Errorf("not repairable")
	})
}
func Merge(a, b model.Record) model.Record {
	if b.UpdatedAt.After(a.UpdatedAt) {
		return b
	}
	return a
}
