/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

import "encoding/xml"

// All the classes that parse XML document into instance of Violation need to implement Parser interface
type Parser interface {
	Parse(Decoder) ([]Violation, error)
}

// All the classes that decode XML document and are used in Parser need to implement Decoder interface
type Decoder interface {
	Token() (t xml.Token, err error)
	DecodeElement(v interface{}, start *xml.StartElement) error
}
