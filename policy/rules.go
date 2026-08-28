package policy

import (
	"strings"
	"training29/model"
)

func NormalizeActor(actor string) string { return strings.TrimSpace(strings.ToLower(actor)) }
func CanEdit(r model.Record, actor string) Decision {
	a := NormalizeActor(actor)
	if a == "" {
		return Decision{false, "actor required"}
	}
	if r.IsFinal() {
		return Decision{false, "record final"}
	}
	return Decision{true, "editable"}
}
func CanView(u model.User, r model.Record) Decision {
	if !u.Active {
		return Decision{false, "inactive"}
	}
	if u.ID == r.UserID || u.Department == "培训" {
		return Decision{true, "visible"}
	}
	return Decision{false, "private"}
}
