package audit

import (
	"fmt"
	"time"
	"training29/model"
	"training29/storage"
)

type Logger struct{ Store *storage.Store }

func New(s *storage.Store) *Logger { return &Logger{Store: s} }
func (l *Logger) Record(record, actor, action, detail string) error {
	a := model.Audit{ID: fmt.Sprintf("audit-%d", time.Now().UnixNano()), RecordID: record, Actor: actor, Action: action, Detail: detail, At: time.Now().UTC()}
	return l.Store.SaveAudit(a)
}
func (l *Logger) Read(id string) (model.Audit, error) { return l.Store.GetAudit(id) }
func (l *Logger) Valid(a model.Audit) bool            { return a.Valid() }
