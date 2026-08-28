package storage

import (
	"path/filepath"
	"testing"
	"training29/model"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := model.NewRecord("r1", "u1", "e1")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("r1"); e != nil {
		t.Fatal(e)
	}
}
