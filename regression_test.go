package main

import (
	"path/filepath"
	"sync"
	"testing"
	"time"
	"training29/enroll"
	"training29/model"
	"training29/storage"
)

func TestRecordFlow29(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	svc := enroll.New(s)
	e := model.Event{ID: "training-29", Capacity: 20, Open: true, StartsAt: time.Now().Add(time.Hour)}
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			u := model.User{ID: "u" + string(rune('a'+i)), Name: "N", Email: "x@y", Active: true}
			_, _ = svc.Register(u, e)
		}(i)
	}
	wg.Wait()
	rs, _ := s.ListRecords()
	if len(rs) != 2 {
		t.Fatalf("concurrent registrations: got %d want 2", len(rs))
	}
}
