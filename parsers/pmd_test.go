// Package parser provides set of classes that helps parse various QA
// tools output into instances of []Violation
package parsers

import (
	"encoding/xml"
	"reflect"
	"testing"
)

type PmdTest struct {
	Counter     int
	Tokens      []xml.Token
	TokensErr   []error
	Elements    []Mess
	ElementsErr []error
}

func (ct *PmdTest) Token() (t xml.Token, err error) {
	ct.Counter = ct.Counter + 1
	return ct.Tokens[ct.Counter], ct.TokensErr[ct.Counter]
}

func (ct *PmdTest) DecodeElement(v interface{}, start *xml.StartElement) error {
	g := reflect.ValueOf(v).Elem()
	gv := reflect.ValueOf(ct.Elements[ct.Counter])
	g.Set(gv)
	return ct.ElementsErr[ct.Counter]
}

func TestNormalPmd(t *testing.T) {
	ct := &PmdTest{
		Counter: -1,
		Tokens: []xml.Token{
			xml.StartElement{
				xml.Name{"", "violation"},
				[]xml.Attr{},
			},
			xml.StartElement{
				xml.Name{"", "donald-duck"},
				[]xml.Attr{},
			},
			xml.StartElement{
				xml.Name{"", "violation"},
				[]xml.Attr{},
			},
			nil,
		},
		TokensErr: []error{nil, nil, nil, nil},
		Elements: []Mess{
			Mess{
				Rule:     "Rule no 1",
				RuleSet:  "Rule set no 1",
				Url:      "http://example.com/1/1.html",
				Priority: 1,
				Message:  "Fake message no 1",
				File: File{
					Name:     "/home/foo/project/bar.go",
					FromLine: 10,
					ToLine:   12,
				},
			},
			{},
			Mess{
				Rule:     "Rule no 2",
				RuleSet:  "Rule set no 2",
				Url:      "http://example.com/2/2.html",
				Priority: 2,
				Message:  "Fake message no 2",
				File: File{
					Name:     "/home/foo/project/bar.go",
					FromLine: 33,
					ToLine:   34,
				},
			},
			{},
		},
		ElementsErr: []error{nil, nil, nil, nil},
	}
	c := new(Pmd)
	v, _ := c.Parse(ct)
	assertViolation(t, v[0], "pmd", 1, "Rule \"Rule no 1\" from set \"Rule set no 1\" has been violated with message: \"Fake message no 1\" (for details see: http://example.com/1/1.html)")
	assertViolation(t, v[1], "pmd", 1, "Rule \"Rule no 1\" from set \"Rule set no 1\" has been violated with message: \"Fake message no 1\" (for details see: http://example.com/1/1.html)")
	assertViolation(t, v[2], "pmd", 1, "Rule \"Rule no 1\" from set \"Rule set no 1\" has been violated with message: \"Fake message no 1\" (for details see: http://example.com/1/1.html)")
	assertViolation(t, v[3], "pmd", 2, "Rule \"Rule no 2\" from set \"Rule set no 2\" has been violated with message: \"Fake message no 2\" (for details see: http://example.com/2/2.html)")
	assertViolation(t, v[4], "pmd", 2, "Rule \"Rule no 2\" from set \"Rule set no 2\" has been violated with message: \"Fake message no 2\" (for details see: http://example.com/2/2.html)")
}
