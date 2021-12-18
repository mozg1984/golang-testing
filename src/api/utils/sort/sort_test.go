package sort

import (
	"testing"
)

// Code coverage is not a good metric
// because it's only shows us how many
// lines of code were executed while doing tests.
// In other words we can just simply run functions without
// validation parts and the code coverage will be high.

func TestBubbleSortIncreasingOrder(t *testing.T) {
	// Init
	elements := []int{8, 2, 5, 1, 0}

	// Execution
	BubbleSort(elements, Ascending)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[4] != 8 {
		t.Error("first element should be 8")
	}
}

func TestBubbleSortDecreasingOrder(t *testing.T) {
	// Init
	elements := []int{8, 2, 5, 1, 0}

	// Execution
	BubbleSort(elements, Descending)

	// Validation
	if elements[0] != 8 {
		t.Error("first element should be 8")
	}
	if elements[4] != 0 {
		t.Error("first element should be 0")
	}
}
