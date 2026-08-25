package report

import (
	"fmt"
	"strings"
	"training29/model"
)

func Header() string { return "培训29报名名单" }
func Row(r model.Record) string {
	return fmt.Sprintf("%s | %s | %s | %s", r.ID, r.UserID, r.EventID, r.Status)
}
func Markdown(rs []model.Record) string {
	lines := []string{"| ID | User | Event | Status |", "|---|---|---|---|"}
	for _, r := range rs {
		lines = append(lines, fmt.Sprintf("| %s | %s | %s | %s |", r.ID, r.UserID, r.EventID, r.Status))
	}
	return strings.Join(lines, "\n")
}
func Group(rs []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range rs {
		out[r.EventID] = append(out[r.EventID], r)
	}
	return out
}
func CountStatus(rs []model.Record, status string) int {
	n := 0
	for _, r := range rs {
		if r.Status == status {
			n++
		}
	}
	return n
}
