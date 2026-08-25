package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"training29/model"
)

func (s *Store) EachRecord(fn func(model.Record) error) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			return fn(r)
		})
	})
}
func (s *Store) CountBucket(name string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(name)).ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
func (s *Store) ClearBucket(name string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(name))
		return b.ForEach(func(k, _ []byte) error { return b.Delete(k) })
	})
}
