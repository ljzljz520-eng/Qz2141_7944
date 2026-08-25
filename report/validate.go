package report

import (
	"fmt"
	"training29/model"
)

func ValidateRecords(rs []model.Record) []error {
	errs := []error{}
	seen := map[string]bool{}
	for _, r := range rs {
		if r.ID == "" {
			errs = append(errs, fmt.Errorf("empty id"))
		}
		if seen[r.ID] {
			errs = append(errs, fmt.Errorf("duplicate %s", r.ID))
		}
		seen[r.ID] = true
		if r.Version < 1 {
			errs = append(errs, fmt.Errorf("bad version %s", r.ID))
		}
		if !r.CanTransition(r.Status) && r.Status != "submitted" && r.Status != "processing" && r.Status != "approved" && r.Status != "rejected" && r.Status != "archived" {
			errs = append(errs, fmt.Errorf("bad status %s", r.Status))
		}
	}
	return errs
}
func Ready(s Summary) bool {
	return s.Total > 0 && s.Submitted+s.Processing+s.Approved+s.Rejected+s.Archived == s.Total
}
