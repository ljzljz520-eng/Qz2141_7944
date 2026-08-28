package main

import (
	"path/filepath"
	"testing"
	"time"
	"training29/catalog"
	"training29/enroll"
	"training29/model"
	"training29/query"
	"training29/storage"
	"training29/workflow"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	c := catalog.New(s)
	e := c.SeedEvent()
	c.SaveEvent(e)
	u := model.User{ID: "u", Name: "User", Email: "u@e", Active: true}
	svc := enroll.New(s)
	if _, x := svc.Register(u, e); x != nil {
		t.Fatal(x)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := model.NewRecord("r", "u", "e")
	s.SaveRecord(r)
	p := workflow.New(s)
	if _, e := p.Process("r", "reviewer"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveRecord(model.NewRecord("r", "u", "e"))
	q := query.New(s)
	rs, e := q.Search("u")
	if e != nil || len(rs) != 1 {
		t.Fatal(e)
	}
}

var _ = time.Now
