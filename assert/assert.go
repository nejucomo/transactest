/* Utilities that are probably / should be in the go stdlib, but I don't
 * know the language well enough and don't have internet access.
 */
package assert

import (
	"fmt"
	"log"
)

func Assert(cond bool, tmpl string, arguments ...interface{}) {
	if !cond {
		msg := fmt.Sprintf(tmpl, arguments...)
		log.Fatalf("Assertion failure: %s\n", msg)
	}
}

func Unreachable(tmpl string, arguments ...interface{}) {
	fail_assertion("unreachable code", tmpl, arguments...)
}

func NotImplemented(tmpl string, arguments ...interface{}) {
	fail_assertion("not implemented", tmpl, arguments...)
}

func fail_assertion(prefix, tmpl string, arguments ...interface{}) {
	Assert(false, fmt.Sprintf("%s: %s", prefix, tmpl), arguments...)
}
