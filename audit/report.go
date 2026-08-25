package audit

import (
	"sort"
	"training29/model"
)

func GroupByAction(as []model.Audit) map[string]int {
	out := map[string]int{}
	for _, a := range as {
		out[a.Action]++
	}
	return out
}
func Recent(as []model.Audit, n int) []model.Audit {
	sort.Slice(as, func(i, j int) bool { return as[i].At.After(as[j].At) })
	if n > len(as) {
		n = len(as)
	}
	return as[:n]
}
func Describe(a model.Audit) string { return a.Actor + ":" + a.Action + ":" + a.RecordID }
