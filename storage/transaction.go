package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"training29/model"
)

func (s *Store) SaveBundle(r model.Record, a model.Audit) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		rb, _ := json.Marshal(r)
		if e := tx.Bucket([]byte("records")).Put([]byte(r.ID), rb); e != nil {
			return e
		}
		ab, _ := json.Marshal(a)
		return tx.Bucket([]byte("audits")).Put([]byte(a.ID), ab)
	})
}
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
func (s *Store) Exists(id string) bool { _, e := s.GetRecord(id); return e == nil }
