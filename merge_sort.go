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

		for k := 0; k < len(collectionSorted); k++ {
			collection[k] = collectionSorted[k]
		}
	}
}

func merge(collectionA []int, collectionB []int) []int {
	na, nb := len(collectionA), len(collectionB)
	collectionSorted := make([]int, na+nb)
	i, j := 0, 0

	for i < na && j < nb {
		if collectionA[i] > collectionB[j] {
			collectionSorted[i+j] = collectionB[j]
			j++
		} else {
			collectionSorted[i+j] = collectionA[i]
			i++
		}
	}

	for ; i < na; i++ {
		collectionSorted[i+j] = collectionA[i]
	}

	for ; j < nb; j++ {
		collectionSorted[i+j] = collectionB[j]
	}

	return collectionSorted
}
