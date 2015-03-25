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
	spec, err := parseTestSpec(data)
	checkErr(err)

	runTestSpec(spec)
}

func parseTestSpec(data []byte) (TestSpec, error) {
	var spec TestSpec

	err := json.Unmarshal(data, &spec)
	return spec, err
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
	Transactions []TransactionAssertions
}

type TransactionAssertions struct {
	Transaction Transaction
	Assertions Assertions
}

type Transaction struct {
	Data []byte
	GasLimit Gas
	GasPrice Gas
	Nonce SeqNum
	Sender AccountId
	To AccountId
	Value Ether
}

type Gas uint
type SeqNum uint
type AccountId uint
type Ether uint

type Assertions struct {
	Accounts map[AccountId]AccountAssertion
}

type AccountAssertion struct {
	Balance Ether
	Code ByteCode
	Nonce SeqNum
	Storage map[string]string
}

type ByteCode []byte
