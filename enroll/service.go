package enroll

import (
	"fmt"
	"training29/model"
	"training29/storage"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) Register(u model.User, e model.Event) (model.Record, error) {
	if !u.Valid() {
		return model.Record{}, fmt.Errorf("invalid user")
	}
	if !e.Available() {
		return model.Record{}, fmt.Errorf("event unavailable")
	}
	r := model.NewRecord(e.ID, u.ID, e.ID)
	return r, s.Store.SaveRecord(r)
}
func (s *Service) Submit(r model.Record) error {
	if r.Status != "submitted" {
		return fmt.Errorf("not submitted")
	}
	return s.Store.SaveRecord(r)
}
func (s *Service) Count() int                           { rs, _ := s.Store.ListRecords(); return len(rs) }
func (s *Service) Find(id string) (model.Record, error) { return s.Store.GetRecord(id) }
func (s *Service) Validate(r model.Record) error {
	if r.ID == "" || r.UserID == "" || r.EventID == "" {
		return fmt.Errorf("missing identity")
	}
	return nil
}
