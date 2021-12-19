package services

import "golang-testing/src/api/utils/sort"

const (
	BUBBLE_SORTING_LIMIT = 20000
)

func AscendingSort(elements []int) {
	if len(elements) <= BUBBLE_SORTING_LIMIT {
		sort.BubbleSort(elements, sort.Ascending)
		return
	}

	sort.Sort(elements)
}

func DescendingSort(elements []int) {
	sort.BubbleSort(elements, sort.Descending)
}
