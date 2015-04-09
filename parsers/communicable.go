/*
Package parser provides set of classes that helps parse various QA
tools output into instances of []Violation
*/
package parsers

import "sync"

// Represents object that communicate via channels sending instances of Violation around
// Used as a trait
type Communicable struct {
	ch chan *Violation
	wg *sync.WaitGroup
}

// Sets not exported chan property
func (chn *Communicable) SetChannel(ch chan *Violation) {
	chn.ch = ch
}

// Sets not exported sync.WaitGroup property
func (chn *Communicable) SetWaitGroup(wg *sync.WaitGroup) {
	chn.wg = wg
}
