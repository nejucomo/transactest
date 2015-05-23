package main

import (
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
					Balance:       Ether(*big.NewInt(10000)),
					ContractState: nil,
				},
			},
			[]TransactionAssertions{
				TransactionAssertions{
					Transaction{
						Data:     []byte{},
						GasLimit: Ether(*big.NewInt(123)),
						GasPrice: Ether(*big.NewInt(3)),
						Nonce:    0,
						Sender:   "alice",
						To:       "bob",
						Value:    Ether(*big.NewInt(42)),
					},
					Assertions{},
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
