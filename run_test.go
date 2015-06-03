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
	var (
		startbal int64 = 12345678
		xferamt  int64 = 42
		gprice   int64 = 3
	)

	txgas := params.TxGas.Int64()

	runTestSpecTest(t, 6, 0,
		TestSpec{
			[]Account{
				Account{
					Id:            "alice",
					Balance:       Ether{big.NewInt(startbal)},
					ContractState: nil,
				},
			},
			[]TransactionAssertions{
				TransactionAssertions{
					Transaction{
						Data:     []byte{},
						GasLimit: Ether{params.TxGas},
						GasPrice: Ether{big.NewInt(gprice)},
						Nonce:    0,
						Sender:   "alice",
						To:       "bob",
						Value:    Ether{big.NewInt(xferamt)},
					},
					Assertions{
						// Bug: What about gas costs?
						map[AccountId]AccountAssertion{
							"alice": AccountAssertion{
								Balance: Ether{big.NewInt(startbal - xferamt - gprice*txgas)},
								Code:    nil,
								Nonce:   1,
								Storage: nil,
							},
							"bob": AccountAssertion{
								Balance: Ether{big.NewInt(xferamt)},
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

	results.Print(true, "self test")

	if s != successes || f != failures {
		t.Errorf(
			"Expected %+v successes, found %+v.\nExpected %+v failures, found %+v.\n",
			successes, s,
			failures, f)
		return
	}
}
