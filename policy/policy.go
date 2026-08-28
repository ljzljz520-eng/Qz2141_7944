package policy

import (
	"fmt"
	"training29/model"
)

type Decision struct {
	Allowed bool
	Reason  string
}

func CanRegister(u model.User, e model.Event) Decision {
	if !u.Active {
		return Decision{false, "inactive user"}
	}
	if !e.Open {
		return Decision{false, "event closed"}
	}
	if e.Capacity <= 0 {
		return Decision{false, "event full"}
	}
	return Decision{true, "eligible"}
}
func CanApprove(r model.Record, actor string) Decision {
	if actor == "" {
		return Decision{false, "actor required"}
	}
	if r.Status != "processing" {
		return Decision{false, fmt.Sprintf("status %s", r.Status)}
	}
	return Decision{true, "review complete"}
}
func AllowedStatuses() []string {
	return []string{"submitted", "processing", "approved", "rejected", "archived"}
}
