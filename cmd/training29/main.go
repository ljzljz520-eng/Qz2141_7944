package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"training29/api"
	"training29/audit"
	"training29/catalog"
	"training29/enroll"
	"training29/model"
	"training29/query"
	"training29/storage"
	"training29/workflow"
)

func main() {
	path := "training29.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	c := catalog.New(s)
	ev := c.SeedEvent()
	_ = c.SaveEvent(ev)
	u := model.User{ID: "demo", Name: "Demo", Email: "demo@example.com", Department: "培训", Active: true}
	_ = s.SaveUser(u)
	en := enroll.New(s)
	pr := workflow.New(s)
	q := query.New(s)
	_ = audit.New(s)
	_ = en
	_ = pr
	_ = q
	fmt.Println("training29 listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", api.New(en, q, pr).Routes()))
}
