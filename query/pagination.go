package query

import (
	"net/url"
	"strconv"
	"training29/model"
)

type Page struct {
	Items                []model.Record
	Offset, Limit, Total int
}

func Paginate(rs []model.Record, offset, limit int) Page {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	if offset > len(rs) {
		offset = len(rs)
	}
	end := offset + limit
	if end > len(rs) {
		end = len(rs)
	}
	return Page{Items: rs[offset:end], Offset: offset, Limit: limit, Total: len(rs)}
}
func ParsePage(v url.Values) (int, int) {
	o, _ := strconv.Atoi(v.Get("offset"))
	l, _ := strconv.Atoi(v.Get("limit"))
	return o, l
}
func SelectUsers(rs []model.Record, ids map[string]bool) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if ids[r.UserID] {
			out = append(out, r)
		}
	}
	return out
}
func SelectEvents(rs []model.Record, ids map[string]bool) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if ids[r.EventID] {
			out = append(out, r)
		}
	}
	return out
}
