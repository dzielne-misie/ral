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
			xml.StartElement{
				xml.Name{"", "duplication"},
				[]xml.Attr{
					xml.Attr{xml.Name{"lines", ""}, "512"},
					xml.Attr{xml.Name{"tokens", ""}, "1024"},
				},
			},
			nil,
		},
		TokensErr: []error{nil, nil, nil, nil},
		Elements: []Duplication{
			Duplication{Lines: 32, Tokens: 64, Files: []DFile{DFile{Name: "foo.go", FromLine: 1}, DFile{Name: "bar.go", FromLine: 666}}},
			Duplication{Lines: 128, Tokens: 256, Files: []DFile{DFile{Name: "example.go", FromLine: 55}, DFile{Name: "another_example.go", FromLine: 38}}},
			Duplication{Lines: 512, Tokens: 1024, Files: []DFile{DFile{Name: "foo.go", FromLine: 111}, DFile{Name: "another_example.go", FromLine: 222}}},
		},
		ElementsErr: []error{nil, nil, nil},
	}

	c := NewCpd()
	ch, wg := prepareAndRun(c, ct)
	priorities := []int8{1, 1, 1, 1, 1, 1}
	msgs := []string{
		"32 duplicated lines and 64 duplicated tokens from file foo.go line 1",
		"32 duplicated lines and 64 duplicated tokens from file bar.go line 666",
		"128 duplicated lines and 256 duplicated tokens from file example.go line 55",
		"128 duplicated lines and 256 duplicated tokens from file another_example.go line 38",
		"512 duplicated lines and 1024 duplicated tokens from file foo.go line 111",
		"512 duplicated lines and 1024 duplicated tokens from file another_example.go line 222",
	}
	files := []string{"foo.go", "bar.go", "example.go", "another_example.go", "foo.go", "another_example.go"}
	fromLines := []int16{1, 666, 55, 38, 111, 222}
	toLines := []int16{32, 697, 182, 165, 622, 733}
	go assertViolations(ch, t, "cpd", priorities, msgs, files, fromLines, toLines)
	wg.Wait()
}

// Receives data from channels and performs assertions
func assertViolations(ch chan *Violation, t *testing.T, tp string, priorities []int8, msgs []string, files []string, fromLines []int16, toLines []int16) {
	i := 0
	for {
		select {
		case v := <-ch:
			//			fmt.Println(v)
			assertViolation(t, v, tp, priorities[i], msgs[i], fromLines[i], toLines[i])
			assertFile(t, v.File, files[i])
			i++
		}
	}
}

// Prepares instance of Parser and runs Parse method with mocked dependencies
func prepareAndRun(p Parser, ct Decoder) (ch chan *Violation, wg *sync.WaitGroup) {
	ch = make(chan *Violation, 100)
	wg = new(sync.WaitGroup)
	p.SetChannel(ch)
	p.SetWaitGroup(wg)
	wg.Add(1)
	go p.Parse(ct)
	return ch, wg
}

//assertFile streamlines assertions related to File struct
func assertFile(t *testing.T, f *File, name string) {
	if f.Name != name {
		t.Errorf("Expected File.Name %q. Received - %q", name, f.Name)
	}
}

//assertViolation streamlines assertions related to Violation struct
func assertViolation(t *testing.T, v *Violation, vType string, priority int8, message string, fromLine int16, toLine int16) {
	if v.Type != vType {
		t.Errorf("Expected Violation.Type %q. Received - %q", vType, v.Type)
	}
	if v.Priority != priority {
		t.Errorf("Expected Violation.Priority %q. Received - %q", message, v.Priority)
	}
	if v.Message != message {
		t.Errorf("Expected Violation.Message %q. Received - %q", message, v.Message)
	}
	if v.FromLine != fromLine {
		t.Errorf("Expected Violation.FromLine %q. Received - %q", fromLine, v.FromLine)
	}
	if v.ToLine != toLine {
		t.Errorf("Expected Violation.ToLine %q. Received - %q", toLine, v.ToLine)
	}
}
