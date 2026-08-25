package workflow

import (
	"fmt"
	"training29/model"
)

type Checklist struct {
	Required  []string
	Completed map[string]bool
}

func NewChecklist(items []string) Checklist {
	return Checklist{Required: items, Completed: map[string]bool{}}
}
func (c *Checklist) Complete(item string) error {
	for _, x := range c.Required {
		if x == item {
			c.Completed[item] = true
			return nil
		}
	}
	return fmt.Errorf("unknown checklist item")
}
func (c Checklist) Ready() bool {
	for _, x := range c.Required {
		if !c.Completed[x] {
			return false
		}
	}
	return true
}
func (c Checklist) Remaining() []string {
	out := []string{}
	for _, x := range c.Required {
		if !c.Completed[x] {
			out = append(out, x)
		}
	}
	return out
}
func ValidTransition(from, to string) bool { return model.Record{Status: from}.CanTransition(to) }
