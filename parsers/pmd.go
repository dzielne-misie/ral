/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

import (
	"encoding/xml"
	"fmt"
	"strings"
	"sync"
)

// Pmd strict represents object that allows to parse mess detector files
type Pmd struct {
	ch chan *Violation
	wg *sync.WaitGroup
}

/*
Parse reads mess detector XML file using Decoder and and builds []Violation.
Expects file elements with violation children to be present in the document .
*/
func (pmd *Pmd) Parse(f Decoder) {
	defer pmd.wg.Done()
	for {
		t, _ := f.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "file" {
				var mF MessedFile
				f.DecodeElement(&mF, &se)
				for _, mess := range mF.Violations {
					violation := new(Violation)
					violation.Type = "pmd"
					violation.Priority = mess.Priority
					violation.Message = fmt.Sprintf("Rule %q from set %q has been violated with message: %q (for details see: %s)", mess.Rule, mess.RuleSet, strings.Trim(mess.Message, " \n\t"), mess.Url)
					violation.File.Name = mF.Name
					violation.File.FromLine = mess.FromLine
					violation.File.ToLine = mess.ToLine
					pmd.ch <- violation
				}
			}
		}
	}
}
