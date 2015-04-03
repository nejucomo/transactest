package main

import (
	"github.com/ethereum/go-ethereum/state"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

func runPath(p string) {
	f, err := os.Open(p)
	checkErr(err)
	runReader(f)
}

func runReader(r io.Reader) {
	data, err := ioutil.ReadAll(r)
	checkErr(err)
	runBytes(data)
}

func runBytes(data []byte) {
	spec, err := parseTestSpec(data)
	checkErr(err)

	runTestSpec(spec)
}

func runTestSpec(spec TestSpec) (successes uint, failures uint, err error) {
	successes = 0
	failures = 0
	err = nil

	var sim TestSim

	sim, err = NewTestSim()
	if err != nil {
		return
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

		var s, f uint

		s, f, err = sim.checkAssertions(&ta.Assertions, result, logs, gasleft)
		if err != nil {
			return
		}

		successes += s
		failures += f
	}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
