// Package parser provides set of classes that helps parse various QA
// tools output into instances of violations.Violation
// mess detector etc).
package parsers

import "encoding/xml"
import "github.com/dzielne-misie/ral/violations"

type Parser interface {
	Parse(xml.Decoder) ([]violations.Violation, error)
}

type Decoder interface {
	Token() (t xml.Token, err error)
	DecodeElement(v interface{}, start *xml.StartElement) error
}
