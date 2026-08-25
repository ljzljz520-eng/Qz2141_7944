package workflow

import (
	"path/filepath"
	"testing"
	"training29/model"
	"training29/storage"
)

func TestProcessorTransitions(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	p := New(s)
	r := model.NewRecord("r", "u", "e")
	s.SaveRecord(r)
	if _, e := p.Process("r", "admin"); e != nil {
		t.Fatal(e)
	}
	r, _ = s.GetRecord("r")
	if r.Status != "approved" {
		t.Fatal(r.Status)
	}
}
