package test

import (
	"goalgorithms"
	"testing"
)

func TestMergeSort(t *testing.T) {
	myAlgorithms := goalgorithms.Algorithms{}
	cases := []struct {
		unsortedNumbers []int
		toSortNumbers   []int
		want            []int
	}{
		{
			unsortedNumbers: []int{2, 1, 3, 6, 5},
			toSortNumbers:   []int{2, 1, 3, 6, 5},
			want:            []int{1, 2, 3, 5, 6},
		},
		{
			unsortedNumbers: []int{1, 2, 3, 5, 6},
			toSortNumbers:   []int{1, 2, 3, 5, 6},
			want:            []int{1, 2, 3, 5, 6},
		},
		{
			unsortedNumbers: []int{},
			toSortNumbers:   []int{},
			want:            []int{},
		},
		{
			unsortedNumbers: []int{-5, -2, 3, 5, -10},
			toSortNumbers:   []int{-5, -2, 3, 5, -10},
			want:            []int{-10, -5, -2, 3, 5},
		},
	}

	for _, c := range cases {
		myAlgorithms.MergeSort(c.toSortNumbers)

		if len(c.want) != len(c.toSortNumbers) {
			t.Errorf(
				"len(MergeSort(%v)) == %d, want %d",
				c.unsortedNumbers,
				len(c.unsortedNumbers),
				len(c.want),
			)
		}

		for i := 0; i < len(c.want); i++ {
			if c.want[i] != c.toSortNumbers[i] {
				t.Errorf(
					"MergeSort(%v) == %v, want %v",
					c.unsortedNumbers,
					c.toSortNumbers,
					c.want,
				)

				break
			}
		}
	}
}
