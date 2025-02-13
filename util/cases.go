package util

import (
	"strings"
)

func UpperHead(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:n]) + s[n:]
}
