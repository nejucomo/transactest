package main

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/nejucomo/transactest/report"
	"github.com/nejucomo/transactest/simulator"
	"github.com/nejucomo/transactest/testspec"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

func runPath(p string) (results report.Report, err error) {
	f, err := os.Open(p)
	checkErr(err)
	return runReader(f)
}

func runReader(r io.Reader) (results report.Report, err error) {
	data, err := ioutil.ReadAll(r)
	checkErr(err)
	return runBytes(data)
}

func runBytes(data []byte) (results report.Report, err error) {
	spec, err := testspec.Parse(data)
	checkErr(err)

	return runTestSpec(spec)
}

func runTestSpec(spec testspec.TestSpec) (results report.Report, err error) {
	sim, err2 := simulator.New()
	if err2 != nil {
		err = err2
		return
	}

	for _, acc := range spec.Accounts {
		sim.InitAccount(&acc)
	}

	for _, ta := range spec.Transactions {
		var (
			result  []byte
			logs    state.Logs
			gasleft *big.Int
		)
		result, logs, gasleft, err = sim.ApplyTransaction(&ta.Transaction)
		if err != nil {
			return
		}

		err = sim.CheckAssertions(&results, &ta.Assertions, result, logs, gasleft)
		if err != nil {
			return
		}
	}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
