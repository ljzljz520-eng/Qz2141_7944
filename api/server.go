package api

import (
	"encoding/json"
	"net/http"
	"training29/enroll"
	"training29/query"
	"training29/workflow"
)

type Server struct {
	Enroll   *enroll.Service
	Query    *query.Query
	Workflow *workflow.Processor
}

func New(e *enroll.Service, q *query.Query, w *workflow.Processor) *Server {
	return &Server{Enroll: e, Query: q, Workflow: w}
}
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/records", s.records)
	m.HandleFunc("/records/advance", s.advance)
	return m
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (s *Server) records(w http.ResponseWriter, _ *http.Request) {
	rs, e := s.Query.All()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(rs)
}
func (s *Server) advance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	next := r.URL.Query().Get("status")
	x, e := s.Workflow.Advance(id, next, "api")
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(x)
}
