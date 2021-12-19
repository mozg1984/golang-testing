package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Code coverage is not a good metric
// because it's only shows us how many
// lines of code were executed while doing tests.
// In other words we can just simply run functions without
// validation parts and the code coverage will be high.

func TestBubbleSortIncreasingOrder(t *testing.T) {
	// Init
	elements := []int{8, 2, 5, 1, 0}

	assert.NotNil(t, elements)

	// Execution
	BubbleSort(elements, Ascending)

	// Validation
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 8, elements[len(elements)-1], "last element should be 8")
}

func TestBubbleSortDecreasingOrder(t *testing.T) {
	// Init
	elements := []int{8, 2, 5, 1, 0}

	assert.NotNil(t, elements)

	// Execution
	BubbleSort(elements, Descending)

	// Validation
	assert.EqualValues(t, 8, elements[0], "first element should be 8")
	assert.EqualValues(t, 0, elements[len(elements)-1], "last element should be 0")
}

func TestSortIncreasingOrder(t *testing.T) {
	elements := []int{8, 2, 5, 1, 0}

	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 8, elements[len(elements)-1], "last element should be 8")
}

func BenchmarkBubbleSortIncreasingOrder(b *testing.B) {
	elements := []int{8, 2, 5, 1, 0}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements, Ascending)
	}
}

func BenchmarkSortIncreasingOrder(b *testing.B) {
	elements := []int{8, 2, 5, 1, 0}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
