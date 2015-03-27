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
	err = errors.New("not implemented")
	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
