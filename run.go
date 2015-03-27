package main

import (
	"io"
	"os"
	"log"
	"io/ioutil"
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

	sim := TestSim{} // In the future initialize this from TestSpec.

	for _, ta := range spec.Transactions {
		err2 := sim.applyTransaction(&ta.Transaction)
		if err2 != nil {
			err = err2
			return
		}

		s, f, err3 := sim.checkAssertions(&ta.Assertions)
		if err3 != nil {
			err = err3
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
