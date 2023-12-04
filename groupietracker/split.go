package groupietracker

import (
	"strings"
)

func Split(location string) (string, string) {
	before, after, _ := cut(location, "-")
	before = strings.Title(before)
	after = strings.Title(after)
	return before, after
}

func cut(s, sep string) (string, string, bool) {
	before, after, found := strings.Cut(s, sep)
	return before, after, found
}
