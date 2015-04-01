// Package ral provides set of classes that helps you with
// parsing various QA tools output files (xunit, copy-paste detector,
// mess detector etc).
package parsers

type Violation struct {
	Type     string
	Priority int8
	Message  string
	File     File
}

type File struct {
	Name     string `xml:"path,attr"`
	FromLine int16  `xml:"line,attr"`
	ToLine   int16
}

type Duplication struct {
	Lines  int16  `xml:"lines,attr"`
	Tokens int32  `xml:"tokens,attr"`
	Code   string `xml:"codefragment"`
	Files  []File `xml:"file"`
}

type Mess struct {
	Rule     string
	RuleSet  string
	Url      string
	Priority int8
	Message  string
	File     File
}
