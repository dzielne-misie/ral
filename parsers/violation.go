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

type MessedFile struct {
	Name       string `xml:"name,attr"`
	Violations []Mess `xml:"violation"`
}

type Mess struct {
	Rule     string `xml:"rule,attr"`
	RuleSet  string `xml:"ruleset,attr"`
	Url      string `xml:"externalInfoUrl,attr"`
	Priority int8   `xml:"priority,attr"`
	Message  string `xml:",innerxml"`
	FromLine int16  `xml:"beginline,attr"`
	ToLine   int16  `xml:"endline,attr"`
}
