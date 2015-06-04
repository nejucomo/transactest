package main

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/nejucomo/transactest/testspec"
	"testing"
)

func Test_TestSpec_run_empty(t *testing.T) {
	runTestSpecTest(t, 0, 0,
		testspec.TestSpec{
			[]testspec.Account{},
			[]testspec.TransactionAssertions{},
		},
	)
}

func Test_TestSpec_run_simple_single_xfer(t *testing.T) {
	var (
		startbal int64 = 12345678
		xferamt  int64 = 42
		gprice   int64 = 3
	)

	txgas := params.TxGas.Int64()

	runTestSpecTest(t, 6, 0,
		testspec.TestSpec{
			[]testspec.Account{
				testspec.Account{
					Id:            "alice",
					Balance:       testspec.EtherFromInt64(startbal),
					ContractState: nil,
				},
			},
			[]testspec.TransactionAssertions{
				testspec.TransactionAssertions{
					testspec.Transaction{
						Data:     []byte{},
						GasLimit: testspec.EtherFromBigInt(params.TxGas),
						GasPrice: testspec.EtherFromInt64(gprice),
						Nonce:    0,
						Sender:   "alice",
						To:       "bob",
						Value:    testspec.EtherFromInt64(xferamt),
					},
					testspec.Assertions{
						// Bug: What about gas costs?
						map[testspec.AccountId]testspec.AccountAssertion{
							"alice": testspec.AccountAssertion{
								Balance: testspec.EtherFromInt64(startbal - xferamt - gprice*txgas),
								Code:    nil,
								Nonce:   1,
								Storage: nil,
							},
							"bob": testspec.AccountAssertion{
								Balance: testspec.EtherFromInt64(xferamt),
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

func runTestSpecTest(t *testing.T, successes, failures uint, spec testspec.TestSpec) {
	report, err := runTestSpec(spec)

	if err != nil {
		t.Errorf("Did not expect this error: %+v\n", err)
		return
	}

	s, f := report.Counts()

	report.Print(true, "self test")

	if s != successes || f != failures {
		t.Errorf(
			"Expected %+v successes, found %+v.\nExpected %+v failures, found %+v.\n",
			successes, s,
			failures, f)
		return
	}
}
