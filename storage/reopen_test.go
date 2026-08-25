package storage

import (
	"path/filepath"
	"testing"
	"training29/model"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "persist.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord("persist", "u", "e")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("persist")
	if e != nil || got.ID != "persist" {
		t.Fatalf("%v %#v", e, got)
	}
}
