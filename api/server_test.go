package api

import (
	"net/http/httptest"
	"path/filepath"
	"testing"
	"training29/enroll"
	"training29/query"
	"training29/storage"
	"training29/workflow"
)

func TestHealth(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	h := New(enroll.New(s), query.New(s), workflow.New(s)).Routes()
	r := httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
