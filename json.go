package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
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

// CodeSource can either be hex-encoded bytes or a relative path to compile:
func (self *CodeSource) MarshalJSON() ([]byte, error) {
	var prefix string

	if self.Type == HEX {
		prefix = "hex"
	} else if self.Type == COMPILE {
		prefix = "compile"
	} else {
		self.CheckType()
	}

	return json.Marshal(fmt.Sprintf("%s:%s", prefix, self.Info))
}

func (self *CodeSource) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	parts := strings.SplitN(s, ":", 2)
	if len(parts) == 2 {
		tagname := parts[0]
		self.Info = parts[1]

		if tagname == "hex" {
			self.Type = HEX
		} else if tagname == "compile" {
			self.Type = COMPILE
		} else {
			return errors.New(
				fmt.Sprintf(
					"Expected prefix \"hex:\" or \"compile:\", found: \"%s:\"",
					tagname))
		}

		return nil
	} else {
		return errors.New(fmt.Sprintf("Expected a single ':', found: %#v", s))
	}
}
