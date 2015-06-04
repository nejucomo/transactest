package report

import (
	"fmt"
)

type Report struct {
	rs []result
}

func (self *Report) Print(all bool, label string) {
	fmt.Printf("results for %+v:\n", label)
	for _, r := range self.rs {
		var tag string
		if r.ok {
			tag = "pass"
		} else {
			tag = "FAIL"
		}

		if all || !r.ok {
			fmt.Printf("  %s - %s\n", tag, r.string())
		}
	}
	s, f := self.Counts()
	fmt.Printf("  Successes %+v, Failures %+v\n", s, f)
}

func (self *Report) Record(ok bool, format string, arguments ...interface{}) {
	self.rs = append(self.rs, result{ok, format, arguments})
}

func (self *Report) Counts() (successes, failures uint) {
	for _, r := range self.rs {
		if r.ok {
			successes += 1
		} else {
			failures += 1
		}
	}
	return
}

type result struct {
	ok        bool
	format    string
	arguments []interface{}
}

func (r *result) string() string {
	return fmt.Sprintf(r.format, r.arguments...)
}
