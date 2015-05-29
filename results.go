package main

import (
	"fmt"
)

type Results struct {
	rs []Result
}

func (self *Results) Record(ok bool, format string, arguments ...interface{}) {
	self.rs = append(self.rs, Result{ok, format, arguments})
}

func (self *Results) Counts() (successes, failures uint) {
	for _, r := range self.rs {
		if r.ok {
			successes += 1
		} else {
			failures += 1
		}
	}
	return
}

func (self *Results) Slice() []Result {
	return self.rs
}

type Result struct {
	ok        bool
	format    string
	arguments []interface{}
}

func (r *Result) Ok() bool {
	return r.ok
}

func (r *Result) String() string {
	return fmt.Sprintf(r.format, r.arguments...)
}
