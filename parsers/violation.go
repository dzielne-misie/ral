/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

// Violation struct represents object that store parsed information on violation details
type Violation struct {
	Type     string
	Priority int8
	Message  string
	Name     string
	FromLine int16
	ToLine   int16
	File     File
}

// File struct represents file that violations is found in
type File struct {
	Name string
}

// DFile struct represents file element in copy paste detector XML output file
type DFile struct {
	Name     string `xml:"path,attr"`
	FromLine int16  `xml:"line,attr"`
}

// Duplication struct represents duplication element in copy paste detector XML output file
type Duplication struct {
	Lines  int16   `xml:"lines,attr"`
	Tokens int32   `xml:"tokens,attr"`
	Code   string  `xml:"codefragment"`
	Files  []DFile `xml:"file"`
}

// MessedFile struct represents file element in mess detector output XML file
type MessedFile struct {
	Name       string `xml:"name,attr"`
	Violations []Mess `xml:"violation"`
}

// Mess struct represents violation element in mess detector output XML file
type Mess struct {
	Rule     string `xml:"rule,attr"`
	RuleSet  string `xml:"ruleset,attr"`
	Url      string `xml:"externalInfoUrl,attr"`
	Priority int8   `xml:"priority,attr"`
	Message  string `xml:",innerxml"`
	FromLine int16  `xml:"beginline,attr"`
	ToLine   int16  `xml:"endline,attr"`
}
