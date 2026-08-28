package workflow

import (
	"fmt"
	"training29/model"
	"training29/storage"
)

type Processor struct{ Store *storage.Store }

func New(s *storage.Store) *Processor { return &Processor{Store: s} }
func (p *Processor) Advance(id, next, actor string) (model.Record, error) {
	r, e := p.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if !r.CanTransition(next) {
		return r, fmt.Errorf("invalid transition %s to %s", r.Status, next)
	}
	r.Status = next
	if e = p.Store.SaveRecord(r); e != nil {
		return r, e
	}
	a := model.Audit{ID: id + "-" + next, RecordID: id, Actor: actor, Action: next, Detail: "workflow transition"}
	return r, p.Store.SaveAudit(a)
}
func (p *Processor) Approve(id, actor string) (model.Record, error) {
	return p.Advance(id, "approved", actor)
}
func (p *Processor) Reject(id, actor string) (model.Record, error) {
	return p.Advance(id, "rejected", actor)
}
func (p *Processor) Archive(id, actor string) (model.Record, error) {
	return p.Advance(id, "archived", actor)
}
func (p *Processor) Process(id, actor string) (model.Record, error) {
	if _, e := p.Advance(id, "processing", actor); e != nil {
		return model.Record{}, e
	}
	return p.Approve(id, actor)
}
