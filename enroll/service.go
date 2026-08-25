package enroll

import (
	"fmt"
	"sync/atomic"
	"time"
	"training29/model"
	"training29/storage"
)

type Service struct{ Store *storage.Store }

// recordSeq is a process-wide monotonic counter used to derive record IDs.
// The counter guarantees uniqueness even when two goroutines register within
// the same nanosecond — a bare timestamp cannot, which is why concurrent
// saves previously collapsed onto a single key.
var recordSeq atomic.Uint64

func New(s *storage.Store) *Service { return &Service{Store: s} }

// nextRecordID mints a process-unique record identifier. The atomic sequence
// differentiates concurrent callers, while the timestamp keeps IDs unique
// across process restarts.
func nextRecordID() string {
	return fmt.Sprintf("rec-%d-%d", recordSeq.Add(1), time.Now().UTC().UnixNano())
}

func (s *Service) Register(u model.User, e model.Event) (model.Record, error) {
	if !u.Valid() {
		return model.Record{}, fmt.Errorf("invalid user")
	}
	if !e.Available() {
		return model.Record{}, fmt.Errorf("event unavailable")
	}
	r := model.NewRecord(nextRecordID(), u.ID, e.ID)
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
