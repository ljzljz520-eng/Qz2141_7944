package catalog

import (
	"fmt"
	"training29/model"
)

func Reserve(e *model.Event) error {
	if e == nil {
		return fmt.Errorf("nil event")
	}
	if !e.Open {
		return fmt.Errorf("closed")
	}
	if e.Capacity <= 0 {
		return fmt.Errorf("full")
	}
	e.Capacity--
	return nil
}
func Release(e *model.Event) {
	if e != nil {
		e.Capacity++
	}
}
func IsPopular(e model.Event) bool { return e.Capacity < 10 }
