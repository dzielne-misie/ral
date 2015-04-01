// Package parser provides set of classes that helps parse various QA
// tools output into instances of Violation
// mess detector etc).
package parsers

import "encoding/xml"

type Parser interface {
	Parse(Decoder) ([]Violation, error)
}

type Decoder interface {
	Token() (t xml.Token, err error)
	DecodeElement(v interface{}, start *xml.StartElement) error
}
