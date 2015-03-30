// Package parser provides set of classes that helps parse various QA
// tools output into instances of []violations.Violation
package parsers

import "testing"
import "encoding/xml"
import "github.com/dzielne-misie/ral/violations"
import "fmt"

type CpdTest struct {
	Counter     int
	Tokens      []xml.Token
	TokensErr   []error
	Elements    []violations.Duplication
	ElementsErr []error
}

func (ct CpdTest) Token() (t xml.Token, err error) {
	ct.Counter = ct.Counter + 1
	return ct.Tokens[ct.Counter], ct.TokensErr[ct.Counter]
}

func (ct CpdTest) DecodeElement(v interface{}, start *xml.StartElement) error {
	return ct.ElementsErr[ct.Counter]
}

func TestNormal(t *testing.T) {
	ct := &CpdTest{
		Counter: -1,
		Tokens: []xml.Token{
			xml.StartElement{
				xml.Name{"duplication", ""},
				[]xml.Attr{
					xml.Attr{xml.Name{"lines", ""}, "32"},
					xml.Attr{xml.Name{"tokens", ""}, "64"},
				},
			},
			xml.StartElement{
				xml.Name{"duplication", ""},
				[]xml.Attr{
					xml.Attr{xml.Name{"lines", ""}, "128"},
					xml.Attr{xml.Name{"tokens", ""}, "256"},
				},
			},
			nil,
		},
		TokensErr: []error{nil, nil},
		Elements: []violations.Duplication{
			violations.Duplication{Lines: 32, Tokens: 64, CopiedFrom: violations.File{Name: "foo.go", Line: 1}, PastedTo: violations.File{Name: "bar.go", Line: 666}},
			violations.Duplication{Lines: 128, Tokens: 44, CopiedFrom: violations.File{Name: "example.go", Line: 55}, PastedTo: violations.File{Name: "another_example.go", Line: 38}},
		},
		ElementsErr: []error{nil, nil},
	}
	c := new(Cpd)
	v, _ := c.Parse(ct)
	assertViolation(t, v[0], "cpd", 1, "32 duplicated lines and 64 duplicated tokens in foo.go")
	fmt.Println(len(v))
}

func assertViolation(t *testing.T, v violations.Violation, vType string, priority int8, message string) {
	if v.Type != vType {
		t.Errorf("Expected Violation.Type %q. Received - %q", vType, v.Type)
	}
}
