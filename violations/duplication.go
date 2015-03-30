// Package ral provides set of classes that helps you with
// parsing various QA tools output files (xunit, copy-paste detector,
// mess detector etc).
package violations

type File struct {
	Name string
	Line int16
}

type Duplication struct {
	Lines      int16
	Tokens     int32
	code       string
	CopiedFrom File
	PastedTo   File
}
