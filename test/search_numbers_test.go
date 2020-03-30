package test

import (
	"goalgorithms"
	"testing"
)

func TestSearchNumbers(t *testing.T) {
	var (
		myAlgorithms = goalgorithms.Algorithms{}
		cases        = []struct {
			numbers []int
			key     int
			want    int
		}{
			{numbers: []int{3, 12, 8, 6}, key: 12, want: 1},
			{numbers: []int{3, 12, 8, 6}, key: 8, want: 2},
			{numbers: []int{3, 12, 8, 6}, key: 25, want: -1},
			{numbers: []int{}, key: 3, want: -1},
		}
	)

	for _, c := range cases {
		got := myAlgorithms.SearchNumbers(c.numbers, c.key)

		if c.want != got {
			t.Errorf("SearchNumbers(%v, %d) == %d, want %d", c.numbers, c.key, got, c.want)
		}
	}
}
