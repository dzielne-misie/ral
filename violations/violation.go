// Package ral provides set of classes that helps you with
// parsing various QA tools output files (xunit, copy-paste detector,
// mess detector etc).
package violations

type Violation struct {
	Type     string
	Priority int8
	Message  string
}
