package unit

import "testing"

const errorMessage = "Value expected %v, result received was %v"

func TestMedian(t *testing.T) {
	expectedValue := 7.28
	value := Median(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(errorMessage, expectedValue, value)
	}
}
