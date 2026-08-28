package enroll

import (
	"path/filepath"
	"testing"
	"time"
	"training29/model"
	"training29/storage"
)

func TestRegisterValidation(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	svc := New(s)
	_, e := svc.Register(model.User{}, model.Event{})
	if e == nil {
		t.Fatal("expected validation")
	}
	u := model.User{ID: "u", Name: "N", Email: "n@e", Active: true}
	ev := model.Event{ID: "e", Capacity: 1, Open: true, StartsAt: time.Now().Add(time.Hour)}
	if _, e = svc.Register(u, ev); e != nil {
		t.Fatal(e)
	}
}
