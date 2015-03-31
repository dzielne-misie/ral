// Package parser provides set of classes that helps parse various QA
// tools output into instances of []violations.Violation
package parsers

import "testing"
import "encoding/xml"
import "github.com/dzielne-misie/ral/violations"
import "reflect"

type CpdTest struct {
	Counter     int
	Tokens      []xml.Token
	TokensErr   []error
	Elements    []violations.Duplication
	ElementsErr []error
}

func (ct *CpdTest) Token() (t xml.Token, err error) {
	ct.Counter = ct.Counter + 1
	return ct.Tokens[ct.Counter], ct.TokensErr[ct.Counter]
}

func (ct *CpdTest) DecodeElement(v interface{}, start *xml.StartElement) error {
	g := reflect.ValueOf(v).Elem()
	gv := reflect.ValueOf(ct.Elements[ct.Counter])
	g.Set(gv)
	return ct.ElementsErr[ct.Counter]
}

func TestNormal(t *testing.T) {
	ct := &CpdTest{
		Counter: -1,
		Tokens: []xml.Token{
			xml.StartElement{
				xml.Name{"", "duplication"},
				[]xml.Attr{
					xml.Attr{xml.Name{"lines", ""}, "32"},
					xml.Attr{xml.Name{"tokens", ""}, "64"},
				},
			},
			xml.StartElement{
				xml.Name{"", "duplication"},
				[]xml.Attr{
					xml.Attr{xml.Name{"lines", ""}, "128"},
					xml.Attr{xml.Name{"tokens", ""}, "256"},
				},
			},
			nil,
		},
		TokensErr: []error{nil, nil, nil},
		Elements: []violations.Duplication{
			violations.Duplication{Lines: 32, Tokens: 64, CopiedFrom: violations.File{Name: "foo.go", FromLine: 1, ToLine: 0}, PastedTo: violations.File{Name: "bar.go", FromLine: 666, ToLine: 0}},
			violations.Duplication{Lines: 128, Tokens: 256, CopiedFrom: violations.File{Name: "example.go", FromLine: 55, ToLine: 0}, PastedTo: violations.File{Name: "another_example.go", FromLine: 38, ToLine: 0}},
		},
		ElementsErr: []error{nil, nil},
	}
	c := new(Cpd)
	v, _ := c.Parse(ct)
	assertViolation(t, v[0], "cpd", 1, "32 duplicated lines and 64 duplicated tokens from file foo.go line 1")
	assertViolation(t, v[1], "cpd", 1, "128 duplicated lines and 256 duplicated tokens from file example.go line 55")
}

func assertViolation(t *testing.T, v violations.Violation, vType string, priority int8, message string) {
	if v.Type != vType {
		t.Errorf("Expected Violation.Type %q. Received - %q", vType, v.Type)
	}
	if v.Priority != priority {
		t.Errorf("Expected Violation.Priority %q. Received - %q", message, v.Priority)
	}
	if v.Message != message {
		t.Errorf("Expected Violation.Message %q. Received - %q", message, v.Message)
	}
}
