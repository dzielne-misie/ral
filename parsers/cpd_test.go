/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

import (
	"encoding/xml"
	"reflect"
	"sync"
	"testing"
)

// PmdTest struct allows us to mock xml.Decoder behaviour.
type CpdTest struct {
	Counter     int
	Tokens      []xml.Token
	TokensErr   []error
	Elements    []Duplication
	ElementsErr []error
}

// Implements Decoder interface for testing purposes
func (ct *CpdTest) Token() (t xml.Token, err error) {
	ct.Counter = ct.Counter + 1
	return ct.Tokens[ct.Counter], ct.TokensErr[ct.Counter]
}

// Implements Decoder interface for testing purposes
func (ct *CpdTest) DecodeElement(v interface{}, start *xml.StartElement) error {
	g := reflect.ValueOf(v).Elem()
	gv := reflect.ValueOf(ct.Elements[ct.Counter])
	g.Set(gv)
	return ct.ElementsErr[ct.Counter]
}

/*
Copy-paste detector parser test.
Tests absolutely normal program execution. No alarms an no suprises.
*/
func TestNormalCpd(t *testing.T) {
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
		Elements: []Duplication{
			Duplication{Lines: 32, Tokens: 64, Files: []File{File{Name: "foo.go", FromLine: 1, ToLine: 0}, File{Name: "bar.go", FromLine: 666, ToLine: 0}}},
			Duplication{Lines: 128, Tokens: 256, Files: []File{File{Name: "example.go", FromLine: 55, ToLine: 0}, File{Name: "another_example.go", FromLine: 38, ToLine: 0}}},
		},
		ElementsErr: []error{nil, nil},
	}
	ch := make(chan *Violation, 100)
	wg := new(sync.WaitGroup)
	c := &Cpd{ch, wg}
	wg.Add(1)
	go c.Parse(ct)
	priorities := []int8{1, 1, 1, 1}
	msgs := []string{
		"32 duplicated lines and 64 duplicated tokens from file foo.go line 1",
		"32 duplicated lines and 64 duplicated tokens from file bar.go line 666",
		"128 duplicated lines and 256 duplicated tokens from file example.go line 55",
		"128 duplicated lines and 256 duplicated tokens from file another_example.go line 38",
	}
	files := []string{"foo.go", "bar.go", "example.go", "another_example.go"}
	fromLines := []int16{1, 666, 55, 38}
	toLines := []int16{32, 697, 182, 165}
	i := 0
	go func() {
		for {
			select {
			case v := <-ch:
				assertViolation(t, v, "cpd", priorities[i], msgs[i])
				assertFile(t, &v.File, files[i], fromLines[i], toLines[i])
				i++
			}
		}
	}()
}

//assertFile streamlines assertions related to File struct
func assertFile(t *testing.T, f *File, name string, fromLine int16, toLine int16) {
	if f.Name != name {
		t.Errorf("Expected File.Name %q. Received - %q", name, f.Name)
	}
	if f.FromLine != fromLine {
		t.Errorf("Expected File.FromLine %d. Received - %d", fromLine, f.FromLine)
	}
	if f.ToLine != toLine {
		t.Errorf("Expected File.ToLine %d. Received - %d", toLine, f.ToLine)
	}
}

//assertViolation streamlines assertions related to Violation struct
func assertViolation(t *testing.T, v *Violation, vType string, priority int8, message string) {
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
