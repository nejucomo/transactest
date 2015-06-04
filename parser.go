package main

import (
	"encoding/json"
	"github.com/nejucomo/transactest/testspec"
)

func parseTestSpec(data []byte) (testspec.TestSpec, error) {
	var spec testspec.TestSpec

	err := json.Unmarshal(data, &spec)
	return spec, err
}
