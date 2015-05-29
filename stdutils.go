/* Utilities that are probably / should be in the go stdlib, but I don't
 * know the language well enough and don't have internet access.
 */
package main

import (
	"math/big"
)

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

func addBigInts(a, b *big.Int) *big.Int {
	/* This is terrible in a few ways:
		 * - Can't find a functional api for addition, so we're writing
		 *   this one.
	     * - Inefficient.
	     * - Pointer API impedance mismatch.
	*/
	sum := big.NewInt(0)
	sum.Add(a, b)
	return sum
}
