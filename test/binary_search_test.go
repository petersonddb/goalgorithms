package test

import (
	"algorithms"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	myAlgorithms := algorithms.Algorithms{}
	cases := []struct {
		numbers []int
		key     int
		want    int
	}{
		{
			numbers: []int{1, 2, 6, 8, 12},
			key:     8,
			want:    3,
		},
		{
			numbers: []int{1, 2, 6, 8, 12},
			key:     2,
			want:    1,
		},
		{
			numbers: []int{1, 2, 6, 8, 12},
			key:     13,
			want:    -1,
		},
		{
			numbers: []int{},
			key:     0,
			want:    -1,
		},
		{
			numbers: []int{-6, -2, -1, 0, 2, 4},
			key:     -2,
			want:    1,
		},
		{
			numbers: []int{-6, -2, -1, 0, 2, 4},
			key:     -10,
			want:    -1,
		},
	}

	for _, c := range cases {
		got := myAlgorithms.BinarySearch(c.numbers, c.key)

		if c.want != got {
			t.Errorf("BinarySearch(%v, %d) == %d, want %d",
				c.numbers,
				c.key,
				got,
				c.want,
			)
		}
	}
}
