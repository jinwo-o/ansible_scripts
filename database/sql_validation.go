package database

import (
	"strings"
	text "text/template"
)

func checkQuery(s string) string {
	return text.HTMLEscapeString(s)
}

func escapeChars(s string) string {
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "%", "\\%", -1)
	s = strings.Replace(s, ",", "\\,", -1)
	s = strings.Replace(s, "'", "\\'", -1)
	s = strings.Replace(s, "-", "\\-", -1)
	s = strings.Replace(s, ";", "\\;", -1)
	s = strings.Replace(s, ";", "\\:", -1)
	s = strings.Replace(s, "|", "\\|", -1)
	s = strings.Replace(s, ">", "\\>", -1)
	s = strings.Replace(s, "<", "\\<", -1)

	return s
}
