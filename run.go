package main

import (
	"github.com/ethereum/go-ethereum/core/state"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

func runPath(p string) (results Results, err error) {
	f, err := os.Open(p)
	checkErr(err)
	return runReader(f)
}

func runReader(r io.Reader) (results Results, err error) {
	data, err := ioutil.ReadAll(r)
	checkErr(err)
	return runBytes(data)
}

func runBytes(data []byte) (results Results, err error) {
	spec, err := parseTestSpec(data)
	checkErr(err)

	return runTestSpec(spec)
}

func runTestSpec(spec TestSpec) (results Results, err error) {
	var sim TestSim

	sim, err = NewTestSim()
	if err != nil {
		return
	}

	for _, acc := range spec.Accounts {
		sim.initAccount(&acc)
	}

	for _, ta := range spec.Transactions {
		var (
			result  []byte
			logs    state.Logs
			gasleft *big.Int
		)
		result, logs, gasleft, err = sim.applyTransaction(&ta.Transaction)
		if err != nil {
			return
		}

		sim.checkAssertions(&results, &ta.Assertions, result, logs, gasleft)
	}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
