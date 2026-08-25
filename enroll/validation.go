package enroll

import (
	"fmt"
	"net/mail"
	"training29/model"
)

func ValidateUser(u model.User) error {
	if !u.Valid() {
		return fmt.Errorf("user fields incomplete")
	}
	if _, e := mail.ParseAddress(u.Email); e != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}
func ValidateEvent(e model.Event) error {
	if e.ID == "" || e.Code == "" {
		return fmt.Errorf("event identity missing")
	}
	if e.Capacity < 1 {
		return fmt.Errorf("capacity must be positive")
	}
	return nil
}
func NormalizeStatus(s string) string {
	switch s {
	case "pending", "new":
		return "submitted"
	case "done":
		return "approved"
	default:
		return s
	}
}
