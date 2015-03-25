package main

import (
	"encoding/json"
)


func parseTestSpec(data []byte) (TestSpec, error) {
	var spec TestSpec

	err := json.Unmarshal(data, &spec)
	return spec, err
}
