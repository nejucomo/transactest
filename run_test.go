package main

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func Test_TestSpec_run_empty(t *testing.T) {
	runTestSpecTest(t, 0, 0, TestSpec{[]Account{}, []TransactionAssertions{}})
}

func Test_TestSpec_run_simple_single_xfer(t *testing.T) {
	runTestSpecTest(t, 0, 0,
		TestSpec{
			[]Account{
				Account{
					Id:            "alice",
					Balance:       Ether(*big.NewInt(12345678)),
					ContractState: nil,
				},
			},
			[]TransactionAssertions{
				TransactionAssertions{
					Transaction{
						Data:     []byte{},
						GasLimit: Ether(*addBigInts(big.NewInt(1234), params.TxGas)),
						GasPrice: Ether(*big.NewInt(3)),
						Nonce:    0,
						Sender:   "alice",
						To:       "bob",
						Value:    Ether(*big.NewInt(42)),
					},
					Assertions{
						// Bug: What about gas costs?
						map[AccountId]AccountAssertion{
							"alice": AccountAssertion{
								Balance: Ether(*big.NewInt(12345678 - 42)),
								Code:    nil,
								Nonce:   1,
								Storage: nil,
							},
							"bob": AccountAssertion{
								Balance: Ether(*big.NewInt(42)),
								Code:    nil,
								Nonce:   0,
								Storage: nil,
							},
						},
					},
				},
			},
		})
}

func runTestSpecTest(t *testing.T, successes, failures uint, spec TestSpec) {
	s, f, err := runTestSpec(spec)

	if err != nil {
		t.Errorf("Did not expect this error: %+v\n", err)
		return
	}

	if s != successes || f != failures {
		t.Errorf(
			"Expected %+v successes, found %+v.\nExpected %+v failures, found %+v.\n",
			successes, s,
			failures, f)
		return
	}
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
