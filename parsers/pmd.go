// Package parser provides set of classes that helps parse various QA
// tools output into instances of []violations.Violation
package parsers

import (
	"encoding/xml"
	"fmt"
	"github.com/dzielne-misie/ral/violations"
)

type Pmd struct{}

func (pmd *Pmd) Parse(f Decoder) (v []violations.Violation, err error) {
	v = make([]violations.Violation, 0, 500)
	err = nil

	for {
		t, _ := f.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "violation" {
				var mess violations.Mess
				f.DecodeElement(&mess, &se)
				for i := mess.File.FromLine; i <= mess.File.ToLine; i++ {
					violation := new(violations.Violation)
					violation.Type = "pmd"
					violation.Priority = mess.Priority
					violation.Message = fmt.Sprintf("Rule %q from set %q has been violated with message: %q (for details see: %s)", mess.Rule, mess.RuleSet, mess.Message, mess.Url)
					v = append(v, *violation)
				}
			}
		}
	}

	return v, err
}
