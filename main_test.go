package main

import "testing"


func Test_parseTestSpec_invalid_json(t *testing.T) {
	parseTestSpec_expectErr(t, "!")
}


func Test_parseTestSpec_malformed_TestSpec_empty_object(t *testing.T) {
	parseTestSpec_expectErr(t, "{}")
}

// Helper code:
func parseTestSpecString(src string) (TestSpec, error) {
	return parseTestSpec([]byte(src))
}

func parseTestSpec_expectErr(t *testing.T, src string) {
	spec, err := parseTestSpec([]byte(src))
	if err == nil {
		t.Errorf("Expected err != nil; found spec == %+v\n", spec)
	}
}

