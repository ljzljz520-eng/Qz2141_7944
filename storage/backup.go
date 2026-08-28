package storage

import (
	"go.etcd.io/bbolt"
	"io"
	"os"
	"path/filepath"
)

func Backup(s *Store, path string) error {
	if s == nil {
		return os.ErrInvalid
	}
	return s.db.View(func(tx *bbolt.Tx) error { return tx.CopyFile(path, 0600) })
}
func Restore(src, dst string) error {
	in, e := os.Open(src)
	if e != nil {
		return e
	}
	defer in.Close()
	out, e := os.Create(filepath.Clean(dst))
	if e != nil {
		return e
	}
	defer out.Close()
	_, e = io.Copy(out, in)
	return e
}
func Healthy(s *Store) bool { return s != nil && s.db != nil }
