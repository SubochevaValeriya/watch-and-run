package model

import "regexp"

type Directory struct {
	Path          string
	Commands      []string
	IncludeRegexp []*regexp.Regexp
	ExcludeRegexp []*regexp.Regexp
	LogFile       string
}
