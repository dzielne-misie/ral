// Package parser provides set of classes that helps parse various QA
// tools output into instances of []Violation
package parsers

import "encoding/xml"
import "fmt"

type Cpd struct {
}

func (cpd *Cpd) Parse(f Decoder) (v []Violation, err error) {
	v = make([]Violation, 0, 500)
	err = nil

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
					violation.File.Name = f.Name
					violation.File.FromLine = f.FromLine
					violation.File.ToLine = f.FromLine + dup.Lines - 1
					v = append(v, *violation)
				}
			}
		}
	}

	return v, err
}
