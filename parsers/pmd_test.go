/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

import (
	"encoding/xml"
	"reflect"
	"testing"
)

// PmdTest struct allows us to mock xml.Decoder behaviour.
type PmdTest struct {
	Counter     int
	Tokens      []xml.Token
	TokensErr   []error
	Elements    []MessedFile
	ElementsErr []error
}

// Implements Decoder interface for testing purposes
func (ct *PmdTest) Token() (t xml.Token, err error) {
	ct.Counter = ct.Counter + 1
	return ct.Tokens[ct.Counter], ct.TokensErr[ct.Counter]

}

// Implements Decoder interface for testing purposes
func (ct *PmdTest) DecodeElement(v interface{}, start *xml.StartElement) error {
	g := reflect.ValueOf(v).Elem()
	gv := reflect.ValueOf(ct.Elements[ct.Counter])
	g.Set(gv)
	return ct.ElementsErr[ct.Counter]
}

/*
Mess detector parser test.
Tests absolutely normal program execution. No alarms an no suprises.
*/
func TestNormalPmd(t *testing.T) {
	ct := &PmdTest{
		Counter: -1,
		Tokens: []xml.Token{
			xml.StartElement{
				xml.Name{"", "file"},
				[]xml.Attr{},
			},
			xml.StartElement{
				xml.Name{"", "donald-duck"},
				[]xml.Attr{},
			},
			xml.StartElement{
				xml.Name{"", "file"},
				[]xml.Attr{},
			},
			nil,
		},
		TokensErr: []error{nil, nil, nil, nil},
		Elements: []MessedFile{
			MessedFile{
				Name: "/home/foo/project/bar.go",
				Violations: []Mess{
					Mess{
						Rule:     "Rule no 1",
						RuleSet:  "Rule set no 1",
						Url:      "http://example.com/1/1.html",
						Priority: 1,
						Message:  "Fake message no 1",
						FromLine: 10,
						ToLine:   12,
					},
					Mess{
						Rule:     "Rule no 2",
						RuleSet:  "Rule set no 1",
						Url:      "http://example.com/1/2.html",
						Priority: 1,
						Message: "\n\n     Fake message no 1				\n\n",
						FromLine: 35,
						ToLine:   88,
					},
				},
			},
			{},
			MessedFile{
				Name: "/home/foo/project/foo.go",
				Violations: []Mess{
					Mess{
						Rule:     "Rule no 2",
						RuleSet:  "Rule set no 2",
						Url:      "http://example.com/2/2.html",
						Priority: 2,
						Message:  "Fake message no 2",
						FromLine: 33,
						ToLine:   99,
					},
				},
			},
			{},
		},
		ElementsErr: []error{nil, nil, nil, nil},
	}
	c := new(Pmd)
	v, _ := c.Parse(ct)
	assertViolation(t, v[0], "pmd", 1, "Rule \"Rule no 1\" from set \"Rule set no 1\" has been violated with message: \"Fake message no 1\" (for details see: http://example.com/1/1.html)")
	assertViolation(t, v[1], "pmd", 1, "Rule \"Rule no 2\" from set \"Rule set no 1\" has been violated with message: \"Fake message no 1\" (for details see: http://example.com/1/2.html)")
	assertViolation(t, v[2], "pmd", 2, "Rule \"Rule no 2\" from set \"Rule set no 2\" has been violated with message: \"Fake message no 2\" (for details see: http://example.com/2/2.html)")
	assertFile(t, v[0].File, "/home/foo/project/bar.go", 10, 12)
	assertFile(t, v[1].File, "/home/foo/project/bar.go", 35, 88)
	assertFile(t, v[2].File, "/home/foo/project/foo.go", 33, 99)
}
