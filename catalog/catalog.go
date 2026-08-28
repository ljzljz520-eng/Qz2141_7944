package catalog

import (
	"fmt"
	"time"
	"training29/model"
	"training29/storage"
)

type Catalog struct{ Store *storage.Store }

func New(s *storage.Store) *Catalog { return &Catalog{Store: s} }
func (c *Catalog) SeedEvent() model.Event {
	return model.Event{ID: "training-29", Code: "29", Title: "培训29", Location: "主楼", Capacity: 100, StartsAt: time.Now().Add(24 * time.Hour), Open: true}
}
func (c *Catalog) SaveEvent(e model.Event) error {
	if e.Code == "" {
		return fmt.Errorf("missing code")
	}
	return c.Store.SaveEvent(e)
}
func (c *Catalog) Event(id string) (model.Event, error) { return c.Store.GetEvent(id) }
func (c *Catalog) Close(id string) error {
	e, x := c.Event(id)
	if x != nil {
		return x
	}
	e.Open = false
	return c.SaveEvent(e)
}
