// Package ral provides set of classes that helps you with
// parsing various QA tools output files (xunit, copy-paste detector,
// mess detector etc).
package violations

type File struct {
	Name     string
	FromLine int16
	ToLine   int16
}

type Duplication struct {
	Lines      int16
	Tokens     int32
	code       string
	CopiedFrom File
	PastedTo   File
}

type Mess struct {
	Rule     string
	RuleSet  string
	Url      string
	Priority int8
	Message  string
	File     File
}
