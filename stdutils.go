/* Utilities that are probably / should be in the go stdlib, but I don't
 * know the language well enough and don't have internet access.
 */
package main

import (
	"fmt"
	"log"
)

func not_implemented(tmpl string, arguments ...interface{}) {
	msg := fmt.Sprintf(tmpl, arguments...)
	log.Fatalf("NOT IMPLEMENTED: %s\n", msg)
}

func bytesEq(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for _, x := range a {
		for _, y := range b {
			if x != y {
				return false
			}
		}
	}
	return true
}
