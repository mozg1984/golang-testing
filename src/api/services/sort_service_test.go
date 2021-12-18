package services

import (
	"testing"
)

func TestAscendingSort(t *testing.T) {
	elements := []int{8, 2, 5, 1, 0}

	AscendingSort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[4] != 8 {
		t.Error("first element should be 8")
	}
}

func TestDescendingSort(t *testing.T) {
	elements := []int{8, 2, 5, 1, 0}

	DescendingSort(elements)

	if elements[0] != 8 {
		t.Error("first element should be 8")
	}
	if elements[4] != 0 {
		t.Error("first element should be 0")
	}
}
