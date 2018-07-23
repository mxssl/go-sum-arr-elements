package main

import "testing"

func TestSum(t *testing.T) {
	inputArray := []int{20, 10, 10, 30, 20, 5, 5}
	rightResult := 100
	result := sumArray(inputArray)

	if result != rightResult {
		t.Errorf("sumArray func is incorrect. Got: %v, Want: %v",
			result, rightResult)
	}
}
