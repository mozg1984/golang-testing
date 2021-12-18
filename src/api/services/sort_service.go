package services

import "golang-testing/src/api/utils/sort"

func AscendingSort(elements []int) {
	sort.BubbleSort(elements, sort.Ascending)
}

func DescendingSort(elements []int) {
	sort.BubbleSort(elements, sort.Descending)
}
