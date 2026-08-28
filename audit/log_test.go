package audit

import (
	"path/filepath"
	"testing"
	"training29/storage"
)

func TestAuditLog(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	l := New(s)
	if e := l.Record("r", "a", "approve", "ok"); e != nil {
		t.Fatal(e)
	}
}
