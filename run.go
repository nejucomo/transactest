package main

import (
	"io"
	"os"
	"log"
	"errors"
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

	spec.run()
}

func (spec *TestSpec) run() (successes uint, failures uint, err error) {
	successes = 0
	failures = 0

	for _, ta := range spec.Transactions {
		s, f, err2 := ta.run()
		if err2 != nil {
			err = err2
			return
		}
		successes += s
		failures += f
	}

	return
}

func (ta *TransactionAssertions) run() (successes uint, failures uint, err error) {
	err = errors.New("Not Implemented.")
	return
}


func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
