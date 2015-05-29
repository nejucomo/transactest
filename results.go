package main

import (
	"fmt"
)

type Results struct {
	rs []Result
}

func (self *Results) Print(all bool, label string) {
	fmt.Printf("Results for %+v:\n", label)
	for _, r := range self.rs {
		var tag string
		if r.ok {
			tag = "pass"
		} else {
			tag = "FAIL"
		}

		if all || !r.ok {
			fmt.Printf("  %s - %s\n", tag, r.String())
		}
	}
	s, f := self.Counts()
	fmt.Printf("  Successes %+v, Failures %+v\n", s, f)
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
