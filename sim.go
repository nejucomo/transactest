package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/tests/helper"
	"github.com/nejucomo/transactest/assert"
	"github.com/nejucomo/transactest/report"
	"github.com/nejucomo/transactest/testspec"
	"math/big"
)

type TestSim struct { // implements vm.Enviroment
	memdb      *ethdb.MemDatabase
	statedb    *state.StateDB
	noncemap   map[testspec.AccountId]testspec.SeqNum
	blocknum   *big.Int
	blockhash  common.Hash
	time       int64
	depth      int
	difficulty *big.Int
	gasLimit   *big.Int
	gas        *big.Int
}

const (
	COINBASE = "COINBASE"
	ORIGIN   = "ORIGIN"
)

func NewTestSim() (sim TestSim, err error) {
	var memdb *ethdb.MemDatabase

	memdb, err = ethdb.NewMemDatabase()
	if err != nil {
		return
	}

	roothash := common.BytesToHash(nil) // FIXME
	statedb := state.New(roothash, memdb)

	// FIXME: Handle many dummy values better:
	sim = TestSim{
		memdb:      memdb,
		statedb:    statedb,
		noncemap:   map[testspec.AccountId]testspec.SeqNum{},
		blocknum:   big.NewInt(0),
		blockhash:  common.BytesToHash(nil),
		time:       0,
		depth:      0,
		difficulty: params.GenesisDifficulty, // FIXME: allow tests to specify.
		gasLimit:   params.GenesisGasLimit,   // FIXME: allow tests to specify.
		gas:        big.NewInt(9876),         // FIXME: determine purpose of this and fix it.
	}
	return
}

func (sim *TestSim) initAccount(acct *testspec.Account) error {
	addr := sim.getAddress(acct.Id)

	obj := state.NewStateObject(*addr, sim.memdb)
	obj.SetBalance(acct.Balance.AsBigInt())

	if acct.ContractState != nil {
		// BUG: What about Storage?
		code, err := sim.loadSource(acct.ContractState.Code)
		if err != nil {
			return err
		}
		obj.SetCode(code)
	}
	sim.statedb.SetStateObject(obj)

	sim.noncemap[acct.Id] = acct.Nonce
	return nil
}

func (sim *TestSim) applyTransaction(txn *testspec.Transaction) (ret []byte, logs state.Logs, gasLeft *big.Int, err error) {
	snapshot := sim.statedb.Copy()

	// Note these need to update per-block when we add multiblock tests:
	coinbase := sim.statedb.GetOrNewStateObject(*sim.getAddress(COINBASE))
	coinbase.SetGasPool(sim.gasLimit)

	message := helper.NewMessage(
		*sim.getAddress(txn.Sender),
		sim.getAddress(txn.To),
		txn.Data,
		txn.Value.AsBigInt(),
		txn.GasLimit.AsBigInt(),
		txn.GasPrice.AsBigInt(),
		uint64(sim.getSenderNonce(txn.Sender)))

	ret, gasLeft, err = core.ApplyMessage(sim, message, coinbase)
	if core.IsNonceErr(err) || core.IsInvalidTxErr(err) {
		sim.statedb.Set(snapshot)
	}
	sim.statedb.Update() // FIXME: What is this?

	return
}

func (sim *TestSim) checkAssertions(report *report.Report, as *testspec.Assertions, applyresult []byte, logs state.Logs, gasleft *big.Int) error {
	for acct, aa := range as.Accounts {
		stob := sim.statedb.GetOrNewStateObject(*sim.getAddress(acct))

		report.Record(
			aa.Balance.AsBigInt().Cmp(stob.Balance()) == 0,
			"Account %+v - Balance: expected %+v vs actual %+v",
			acct,
			aa.Balance.AsBigInt(),
			stob.Balance())

		/* BUG: if we pre-load all sources, this method wouldn't
		 * have an error condition, and compilation problems
		 * would be detected early.
		 */
		code, err := sim.loadSource(aa.Code)
		if err != nil {
			return err
		}

		report.Record(
			bytes.Compare(code, stob.Code()) == 0,
			"Account %+v - Code: expected %+v vs actual %+v",
			acct,
			aa.Code,
			stob.Code())

		report.Record(
			uint64(aa.Nonce) == stob.Nonce(),
			"Account %+v - Nonce: expected %+v vs actual %+v",
			acct,
			aa.Nonce,
			stob.Nonce())

		/* API BUG: the test specifier may want to express "all
		 * of the following keys must exist with the given values"
		 * or they may want "only the following keys...".
		 */
		for key, expected := range aa.Storage {
			actual, ok := stob.Storage()[key]

			var actualdesc string

			if ok {
				actualdesc = fmt.Sprintf("%+v", ok)
			} else {
				actualdesc = "<missing>"
			}

			report.Record(
				ok && expected.Cmp(actual),
				"Account %+v - Storage %+v: expected %+v vs actual %+v",
				acct,
				key,
				expected,
				actualdesc)
		}
	}

	return nil
}

func (sim *TestSim) getAddress(acct testspec.AccountId) *common.Address {
	addr := common.BytesToAddress(sim.getKeyPair(acct).Address())
	return &addr
}

func (sim *TestSim) getKeyPair(acct testspec.AccountId) *crypto.KeyPair {
	kp, err := crypto.NewKeyPairFromSec(sim.getSecretKey(acct))
	if err != nil {
		panic(err)
	}
	return kp
}

func (sim *TestSim) getSecretKey(acct testspec.AccountId) []byte {
	return crypto.Sha3([]byte(acct))
}

func (sim *TestSim) getSenderNonce(acct testspec.AccountId) testspec.SeqNum {
	nonce, ok := sim.noncemap[acct]
	if !ok {
		panic(fmt.Sprintf("Unknown acct %+v in TestSim.GetSenderNonce", acct))
	}
	return nonce
}

func (_ *TestSim) loadSource(src *testspec.CodeSource) ([]byte, error) {
	if src == nil {
		return nil, nil
	} else if src.Type == testspec.HEX {
		return hex.DecodeString(src.Info)
	} else if src.Type == testspec.COMPILE {
		assert.NotImplemented("source compilation for %#v", src.Info)
		return nil, nil
	} else {
		src.CheckType()
		return nil, nil
	}
}
