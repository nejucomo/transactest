package main

import "testing"


func Test_TestSpec_run_empty(t *testing.T) {
	var spec TestSpec = TestSpec{[]TransactionAssertions{}}

	successes, failures, err := spec.run()

	if err != nil {
		t.Errorf("Did not expect this error: %+v\n", err)
		return
	}

	if successes != 0 {
		t.Errorf("Did not expect this successes value: %+v\n", successes)
		return
	}

	if failures != 0 {
		t.Errorf("Did not expect this failures value: %+v\n", failures)
		return
	}
}

