package main

import (
	"io"
	"os"
	"log"
	"io/ioutil"
)


func main() {
	if len(os.Args) > 1 {
		runPath(os.Args[1])
	} else {
		runReader(os.Stdin)
	}
}

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

func runTestSpec(spec TestSpec) {
	log.Fatalln(spec)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
