package audit

import (
	"sort"
	"training29/model"
)

func Timeline(as []model.Audit, record string) []model.Audit {
	out := []model.Audit{}
	for _, a := range as {
		if a.RecordID == record {
			out = append(out, a)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func Actions(as []model.Audit) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, a := range as {
		if !seen[a.Action] {
			seen[a.Action] = true
			out = append(out, a.Action)
		}
	}
	return out
}
