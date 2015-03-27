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

	spec.runTestSpec()
}

func (spec *TestSpec) runTestSpec() {
	log.Fatalln(spec)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
