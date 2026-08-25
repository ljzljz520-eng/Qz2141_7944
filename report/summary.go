package report

import (
	"fmt"
	"sort"
	"strings"
	"training29/model"
)

type Summary struct {
	Total, Submitted, Processing, Approved, Rejected, Archived int
	ByEvent                                                    map[string]int
}

func Build(rs []model.Record) Summary {
	s := Summary{ByEvent: map[string]int{}}
	for _, r := range rs {
		s.Total++
		s.ByEvent[r.EventID]++
		switch r.Status {
		case "submitted":
			s.Submitted++
		case "processing":
			s.Processing++
		case "approved":
			s.Approved++
		case "rejected":
			s.Rejected++
		case "archived":
			s.Archived++
		}
	}
	return s
}
func (s Summary) CompletionRate() float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Approved+s.Archived) / float64(s.Total)
}
func (s Summary) String() string {
	return fmt.Sprintf("total=%d submitted=%d processing=%d approved=%d rejected=%d archived=%d", s.Total, s.Submitted, s.Processing, s.Approved, s.Rejected, s.Archived)
}
func CSV(rs []model.Record) string {
	rows := []string{"id,user,event,status"}
	for _, r := range rs {
		rows = append(rows, strings.Join([]string{r.ID, r.UserID, r.EventID, r.Status}, ","))
	}
	return strings.Join(rows, "\n")
}
func Statuses(rs []model.Record) []string {
	m := map[string]bool{}
	for _, r := range rs {
		m[r.Status] = true
	}
	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
