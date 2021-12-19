package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAscendingSort(t *testing.T) {
	elements := []int{8, 2, 5, 1, 0}

	assert.NotNil(t, elements)

	AscendingSort(elements)

	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 8, elements[len(elements)-1], "last element should be 8")
}

func TestDescendingSort(t *testing.T) {
	elements := []int{8, 2, 5, 1, 0}

	DescendingSort(elements)

	if elements[0] != 8 {
		t.Error("first element should be 8")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
