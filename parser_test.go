package main

import (
	"math/big"
	"testing"
)

func Test_parseTestSpec_invalid_json(t *testing.T) {
	parseTestSpec_expectErr(t, "!")
}

/* skip
func Test_parseTestSpec_malformed_TestSpec_missing_field(t *testing.T) {
	t.Skip("JSON library makes it cumbersome to implement required fields.")
	parseTestSpec_expectErr(t, "{}")
}
*/

/* skip
func Test_parseTestSpec_malformed_TestSpec_unknown_field(t *testing.T) {
	t.Skip("JSON library makes it cumbersome to implement no-unexpected-fields.")
	parseTestSpec_expectErr(t, "{\"fruit\": \"bananas\"}")
}
*/

func Test_parseTestSpec_Transaction_with_value(t *testing.T) {
	spec := parseTestSpec_expectNoErr(t, `{"transactions": [{"transaction": {"value": 42}}]}`)
	if len(spec.Transactions) != 1 {
		t.Errorf("Wrong number of TransactionAssertions: %+v\n", spec)
		return
	}
	expected := big.NewInt(42)
	value := spec.Transactions[0].Transaction.Value.AsBigInt()
	if expected.Cmp(value) != 0 {
		t.Errorf("Expected transaction Value of %+v, but found %+v\n-in: %+v\n", expected, value, spec)
		return
	}
}

// Helper code:
func parseTestSpec_expectNoErr(t *testing.T, src string) TestSpec {
	spec, err := parseTestSpec([]byte(src))
	if err != nil {
		t.Errorf("Unexpected parse error: %+v\n", err)
	}
	return spec
}

func parseTestSpec_expectErr(t *testing.T, src string) {
	spec, err := parseTestSpec([]byte(src))
	if err == nil {
		t.Errorf("Expected err != nil; found spec == %+v\n", spec)
	}
}
