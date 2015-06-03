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
