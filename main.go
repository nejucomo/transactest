package main

import (
	"io"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
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
	var spec TestSpec

	err := json.Unmarshal(data, &spec)
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

type TestSpec struct {
	transactions []TransactionAssertions
}

type TransactionAssertions struct {
	transaction Transaction
	assertions Assertions
}

type Transaction struct {
	data []byte
	gasLimit Gas
	gasPrice Gas
	nonce SeqNum
	sender AccountId
	to AccountId
	value Ether
}

type Gas uint
type SeqNum uint
type AccountId uint
type Ether uint

type Assertions struct {
	accounts map[AccountId]AccountAssertion
}

type AccountAssertion struct {
	balance Ether
	code ByteCode
	nonce SeqNum
	storage StorageAssertion
}

type ByteCode []byte

type StorageAssertion struct {
	data map[string]string
}
