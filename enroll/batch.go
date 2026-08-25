package enroll

import (
	"fmt"
	"training29/model"
)

type BatchResult struct {
	Records []model.Record
	Errors  []error
}

func (s *Service) RegisterBatch(users []model.User, e model.Event) BatchResult {
	out := BatchResult{Records: []model.Record{}, Errors: []error{}}
	for _, u := range users {
		r, x := s.Register(u, e)
		if x != nil {
			out.Errors = append(out.Errors, x)
		} else {
			out.Records = append(out.Records, r)
		}
	}
	return out
}
func (s *Service) Cancel(id string) error {
	r, e := s.Find(id)
	if e != nil {
		return e
	}
	if r.IsFinal() {
		return fmt.Errorf("final record")
	}
	return s.Store.DeleteRecord(id)
}
func (s *Service) Reopen(id string) error {
	r, e := s.Find(id)
	if e != nil {
		return e
	}
	if r.Status != "rejected" {
		return fmt.Errorf("only rejected")
	}
	r.Status = "submitted"
	return s.Store.SaveRecord(r)
}
