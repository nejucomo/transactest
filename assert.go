/* Utilities that are probably / should be in the go stdlib, but I don't
 * know the language well enough and don't have internet access.
 */
package main

import (
	"fmt"
	"log"
)

func assert(cond bool, tmpl string, arguments ...interface{}) {
	if !cond {
		msg := fmt.Sprintf(tmpl, arguments...)
		log.Fatalf("Assertion failure: %s\n", msg)
	}
}

func unreachable(tmpl string, arguments ...interface{}) {
	fail_assertion("unreachable code", tmpl, arguments...)
}

func not_implemented(tmpl string, arguments ...interface{}) {
	fail_assertion("not implemented", tmpl, arguments...)
}

func fail_assertion(prefix, tmpl string, arguments ...interface{}) {
	assert(false, fmt.Sprintf("%s: %s", prefix, tmpl), arguments...)
}
