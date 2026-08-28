package storage

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"os"
	"sync"
	"training29/model"
)

var buckets = []string{"records", "users", "events", "audits"}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(n)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) put(bucket, key string, v any) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), b) })
}
func (s *Store) get(bucket, key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if b == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(b, v)
	})
}
func (s *Store) SaveRecord(r model.Record) error { return s.put("records", r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.get("records", id, &r)
	return r, e
}
func (s *Store) SaveUser(u model.User) error { return s.put("users", u.ID, u) }
func (s *Store) GetUser(id string) (model.User, error) {
	var u model.User
	e := s.get("users", id, &u)
	return u, e
}
func (s *Store) SaveEvent(v model.Event) error { return s.put("events", v.ID, v) }
func (s *Store) GetEvent(id string) (model.Event, error) {
	var v model.Event
	e := s.get("events", id, &v)
	return v, e
}
func (s *Store) SaveAudit(a model.Audit) error { return s.put("audits", a.ID, a) }
func (s *Store) GetAudit(id string) (model.Audit, error) {
	var a model.Audit
	e := s.get("audits", id, &a)
	return a, e
}
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) UpdateRecord(id string, fn func(*model.Record) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if e = fn(&r); e != nil {
		return e
	}
	r.Version++
	return s.SaveRecord(r)
}

var ErrConflict = errors.New("record conflict")
