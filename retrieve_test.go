package main

import (
	"testing"
)

func TestRetrieve(t *testing.T) {
	result, err := Retrieve()

	// Case: Should not return error
	if err != nil {
		t.Error("There is error!")
	}

	// Case: Function should return as we expect
	if *result != "Raymond" {
		t.Error("Result is not as expected!")
	}
}
