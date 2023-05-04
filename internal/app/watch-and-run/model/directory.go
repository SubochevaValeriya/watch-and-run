package model

type Directory struct {
	Path          string
	Commands      []string
	IncludeRegexp map[string]struct{}
	ExcludeRegexp map[string]struct{}
	LogFile       string
}
