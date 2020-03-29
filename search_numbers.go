package algorithms

// SearchNumbers searches the key position inside a collection of numbers
func (a Algorithms) SearchNumbers(numbers []int, key int) int {
	for i, n := range numbers {
		if n == key {
			return i
		}
	}

	return -1
}
