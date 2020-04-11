package goalgorithms

// MergeSort sorts a collection of numbers
func (a Algorithms) MergeSort(collection []int) {
	n := len(collection)

	if n > 1 {
		var collectionSorted []int
		middle := n / 2
		collectionA := collection[0:middle]
		collectionB := collection[middle:n]

		a.MergeSort(collectionA)
		a.MergeSort(collectionB)

		collectionSorted = merge(collectionA, collectionB)

		copy(collection, collectionSorted)
	}
}

func merge(collectionA []int, collectionB []int) []int {
	na, nb := len(collectionA), len(collectionB)
	n := na + nb
	collectionSorted := make([]int, n)

	for i, j, k := 0, 0, 0; k < n; k++ {
		if i >= na {
			collectionSorted[k] = collectionB[j]
			j++
		} else if j >= nb {
			collectionSorted[k] = collectionA[i]
			i++
		} else if collectionA[i] > collectionB[j] {
			collectionSorted[k] = collectionB[j]
			j++
		} else {
			collectionSorted[k] = collectionA[i]
			i++
		}
	}

	return collectionSorted
}
