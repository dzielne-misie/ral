/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

import (
	"encoding/xml"
	"fmt"
)

// Pmd strict represents object that allows to parse copy paste detector files
type cpd struct {
	Communicable
	files *files
}

func NewCpd() *cpd {
	files := NewFiles()
	cpd := &cpd{files: files}
	return cpd
}

/*
Parse reads copy paste detector XML file using Decoder and and builds []Violation.
Expects duplication elements with file children to be present in the document .
*/
func (cpd *cpd) Parse(f Decoder) {
	defer cpd.wg.Done()
	for {
		t, _ := f.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "duplication" {
				var dup Duplication
				f.DecodeElement(&dup, &se)
				for _, f := range dup.Files {
					violation := new(Violation)
					violation.Type = "cpd"
					violation.Priority = 1
					violation.Message = fmt.Sprintf("%d duplicated lines and %d duplicated tokens from file %s line %d", dup.Lines, dup.Tokens, f.Name, f.FromLine)
					violation.File = cpd.files.Get(f.Name)
					violation.FromLine = f.FromLine
					violation.ToLine = f.FromLine + dup.Lines - 1
					cpd.ch <- violation
				}
			}
		}
	}
}
