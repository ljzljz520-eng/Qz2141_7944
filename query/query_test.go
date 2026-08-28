package query

import (
	"path/filepath"
	"testing"
	"training29/model"
	"training29/storage"
)

func TestQueryFilters(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveRecord(model.NewRecord("r", "u", "e"))
	q := New(s)
	rs, e := q.ByStatus("submitted")
	if e != nil || len(rs) != 1 {
		t.Fatalf("%v %d", e, len(rs))
	}
}
