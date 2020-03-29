package algorithms

// BinarySearch searches the key position inside the numbers collection
func (a Algorithms) BinarySearch(numbers []int, key int) int {
	var (
		first = 0
		last  = len(numbers) - 1
	)

	for first <= last {
		middle := (last + first) / 2

		if key < numbers[middle] {
			last = middle - 1
		} else if key > numbers[middle] {
			first = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
