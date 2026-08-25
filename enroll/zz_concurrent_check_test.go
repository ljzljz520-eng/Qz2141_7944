package enroll

import (
	"path/filepath"
	"sync"
	"testing"
	"time"
	"training29/model"
	"training29/storage"
)

func TestConcurrentRegisterKeepsBoth(t *testing.T) {
	dir := t.TempDir()
	s, e := storage.Open(filepath.Join(dir, "t.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()

	ev := model.Event{ID: "training-29", Code: "29", Title: "培训29", Location: "主楼", Capacity: 100, Open: true, StartsAt: time.Now().Add(24 * time.Hour)}
	users := []model.User{
		{ID: "u1", Name: "Alice", Email: "a@x.com", Active: true},
		{ID: "u2", Name: "Bob", Email: "b@x.com", Active: true},
	}
	svc := New(s)

	var wg sync.WaitGroup
	wg.Add(len(users))
	for _, u := range users {
		u := u
		go func() {
			defer wg.Done()
			if _, err := svc.Register(u, ev); err != nil {
				t.Errorf("register %s: %v", u.ID, err)
			}
		}()
	}
	wg.Wait()

	rs, _ := s.ListRecords()
	if len(rs) != 2 {
		t.Fatalf("expected 2 records after concurrent save, got %d", len(rs))
	}
	ids := map[string]bool{}
	for _, r := range rs {
		if ids[r.ID] {
			t.Fatalf("duplicate record id %s", r.ID)
		}
		ids[r.ID] = true
	}
}
