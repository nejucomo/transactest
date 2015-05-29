package main

import (
	"math/big"
)

// *Ether has custom JSON decoding to/from 64bit values:
func (self *Ether) MarshalJSON() ([]byte, error) {
	return self.AsBigInt().MarshalJSON()
}

func (self *Ether) UnmarshalJSON(data []byte) error {
	var i big.Int

	err := i.UnmarshalJSON(data)
	if err == nil {
		*self = Ether{&i}
	}
	return err
}
