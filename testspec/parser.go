package testspec

import (
	"encoding/json"
)

func Parse(data []byte) (TestSpec, error) {
	var spec TestSpec

	err := json.Unmarshal(data, &spec)
	return spec, err
}
