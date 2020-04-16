package test

import (
	"goalgorithms"
	"math"
	"testing"
)

func TestDijkstraShortestPath(t *testing.T) {
	myAlgorithms := goalgorithms.Algorithms{}
	cases := []struct {
		regionMap [][]int64
		src       int
		want      []int64
	}{
		{
			regionMap: [][]int64{
				{-1, 10, -1, -1}, // 0
				{-1, -1, 5, 15},  // 1
				{-1, -1, -1, 30}, // 2
				{-1, -1, -1, -1}, // 3
			},
			src:  0,
			want: []int64{0, 10, 15, 25},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1, -1}, // 0
				{-1, -1, 5, 15},  // 1
				{-1, -1, -1, 30}, // 2
				{-1, -1, -1, -1}, // 3
			},
			src:  2,
			want: []int64{math.MaxInt64, math.MaxInt64, 0, 30},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1, -1, -1, -1}, // 0
				{-1, -1, -1, 2, -1, 5},   // 1
				{-1, -1, -1, 1, -1, -1},  // 2
				{-1, -1, -1, -1, 1, -1},  // 3
				{-1, -1, 28, -1, -1, 1},  // 4
				{-1, -1, -1, -1, -1, -1}, // 5
			},
			src:  0,
			want: []int64{0, 10, 41, 12, 13, 14},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1, -1, -1, -1}, // 0
				{-1, -1, -1, 2, -1, 5},   // 1
				{-1, -1, -1, 1, -1, -1},  // 2
				{-1, -1, -1, -1, 1, -1},  // 3
				{-1, -1, 28, -1, -1, 1},  // 4
				{-1, -1, -1, -1, -1, -1}, // 5
			},
			src:  3,
			want: []int64{math.MaxInt64, math.MaxInt64, 29, 0, 1, 2},
		},
		{
			regionMap: [][]int64{
				{-1}, // 0
			},
			src:  0,
			want: []int64{0},
		},
		{
			regionMap: [][]int64{
				{10}, // 0
			},
			src:  0,
			want: []int64{0},
		},
		{
			regionMap: [][]int64{
				{-1, 10}, // 0
				{10, -1}, // 0
			},
			src:  0,
			want: []int64{0, 10},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1}, // 0
				{-1, -1, -1}, // 1
				{-1, -1, -1}, // 2
			},
			src:  0,
			want: []int64{0, 10, math.MaxInt64},
		},
	}

	for _, c := range cases {
		got := myAlgorithms.DijkstraShortestPath(c.regionMap, c.src)

		if len(c.want) != len(got) {
			t.Errorf(
				"len(DijkstraShortestPath(%v, %d)) == %d, want %d",
				c.regionMap,
				c.src,
				len(got),
				len(c.want),
			)
		}

		for i, distance := range got {
			if distance != c.want[i] {
				t.Errorf(
					"DijkstraShortestPath(%v, %d) == %v, want %v",
					c.regionMap,
					c.src,
					got,
					c.want,
				)
			}
		}
	}
}
