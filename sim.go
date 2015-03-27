package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethdb"
)


type TestSim struct {
	db *ethdb.MemDatabase
}


func NewTestSim() (sim TestSim, err error) {
	var db *ethdb.MemDatabase

	db, err = ethdb.NewMemDatabase()
	if err == nil {
		sim = TestSim{db}
	}
	return
}


func (sim *TestSim) applyTransaction(txn *Transaction) error {
	return errors.New("Not Implemented.")
}


func (sim *TestSim) checkAssertions(as *Assertions) (successes uint, failures uint, err error) {
	err = errors.New("Not Implemented.")
	return
}






