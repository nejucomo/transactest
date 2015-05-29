package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func Test_TestSpec_run_empty(t *testing.T) {
	runTestSpecTest(t, 0, 0, TestSpec{[]Account{}, []TransactionAssertions{}})
}

func Test_TestSpec_run_simple_single_xfer(t *testing.T) {
	runTestSpecTest(t, 6, 0,
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
	results, err := runTestSpec(spec)

	if err != nil {
		t.Errorf("Did not expect this error: %+v\n", err)
		return
	}

	s, f := results.Counts()

	fmt.Printf("Results: %+v successes and %+v failures\n", s, f)
	for _, r := range results.Slice() {
		var tag string
		if r.Ok() {
			tag = "pass"
		} else {
			tag = "FAIL"
		}

		fmt.Printf("  %s - %s\n", tag, r.String())
	}

	if s != successes || f != failures {
		t.Errorf(
			"Expected %+v successes, found %+v.\nExpected %+v failures, found %+v.\n",
			successes, s,
			failures, f)
		return
	}
}
