// Package parser provides set of classes that helps parse various QA
// tools output into instances of []violations.Violation
package parsers

import "encoding/xml"
import "github.com/dzielne-misie/ral/violations"

type Cpd struct {
}

func (cpd *Cpd) Parse(f Decoder) (v []violations.Violation, err error) {
	v = make([]violations.Violation, 100, 500)
	err = nil

	for {
		t, _ := f.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "duplication" {
				var dup violations.Duplication
				f.DecodeElement(dup, &se)
				violation := new(violations.Violation)
				violation.Type = "cpd"
			}
		}
		break
	}

	return v, err
}
