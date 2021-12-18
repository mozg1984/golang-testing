package sort

type Ordering func(a, b int) bool

func Ascending(a, b int) bool {
	return a > b
}

func Descending(a, b int) bool {
	return a < b
}

func BubbleSort(elements []int, ordering Ordering) {
	keepWorking := true
	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if ordering(elements[i], elements[i+1]) {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
