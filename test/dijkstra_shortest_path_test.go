package test

import (
	"goalgorithms"
	"math"
	"testing"
)

func TestDijkstraShortestPath(t *testing.T) {
	myAlgorithms := goalgorithms.Algorithms{}
	cases := []struct {
		regionMap       [][]int64
		src             int
		dst             int
		wantedDistances []int64
		wantedPrevious  []int
	}{
		{
			regionMap: [][]int64{
				{-1, 10, -1, -1}, // 0
				{-1, -1, 5, 15},  // 1
				{-1, -1, -1, 30}, // 2
				{-1, -1, -1, -1}, // 3
			},
			src:             0,
			dst:             3,
			wantedDistances: []int64{0, 10, 15, 25},
			wantedPrevious:  []int{-1, 0, 1, 1},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1, -1}, // 0
				{-1, -1, 5, 15},  // 1
				{-1, -1, -1, 30}, // 2
				{-1, -1, -1, -1}, // 3
			},
			src:             2,
			dst:             3,
			wantedDistances: []int64{math.MaxInt64, math.MaxInt64, 0, 30},
			wantedPrevious:  []int{-1, -1, -1, 2},
		},
		{
			regionMap: [][]int64{
				{-1, 2, -1, 1},   // 0
				{-1, -1, 1, -1},  // 1
				{-1, -1, -1, 1},  // 2
				{-1, -1, -1, -1}, // 3
			},
			src:             0,
			dst:             3,
			wantedDistances: []int64{0, 2, math.MaxInt64, 1},
			wantedPrevious:  []int{-1, 0, -1, 0},
		},
		{
			regionMap: [][]int64{
				{-1, 8, 10, -1, -1},  // 0
				{-1, -1, -1, 4, -1},  // 1
				{-1, -1, -1, -1, -1}, // 2
				{-1, -1, -1, -1, 1},  // 3
				{-1, -1, 1, -1, -1},  // 4
			},
			src:             0,
			dst:             2,
			wantedDistances: []int64{0, 8, 10, 12, math.MaxInt64},
			wantedPrevious:  []int{-1, 0, 0, 1, -1},
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
			src:             0,
			dst:             5,
			wantedDistances: []int64{0, 10, 41, 12, 13, 14},
			wantedPrevious:  []int{-1, 0, 4, 1, 3, 4},
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
			src:             3,
			dst:             5,
			wantedDistances: []int64{math.MaxInt64, math.MaxInt64, 29, 0, 1, 2},
			wantedPrevious:  []int{-1, -1, 4, -1, 3, 4},
		},
		{
			regionMap: [][]int64{
				{-1}, // 0
			},
			src:             0,
			dst:             0,
			wantedDistances: []int64{0},
			wantedPrevious:  []int{-1},
		},
		{
			regionMap: [][]int64{
				{10}, // 0
			},
			src:             0,
			dst:             0,
			wantedDistances: []int64{0},
			wantedPrevious:  []int{-1},
		},
		{
			regionMap: [][]int64{
				{-1, 10}, // 0
				{10, -1}, // 1
			},
			src:             0,
			dst:             1,
			wantedDistances: []int64{0, 10},
			wantedPrevious:  []int{-1, 0},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1}, // 0
				{-1, -1, -1}, // 1
				{-1, -1, -1}, // 2
			},
			src:             0,
			dst:             1,
			wantedDistances: []int64{0, 10, math.MaxInt64},
			wantedPrevious:  []int{-1, 0, -1},
		},
		{
			regionMap: [][]int64{
				{-1, 10, -1}, // 0
				{-1, -1, -1}, // 1
				{-1, -1, -1}, // 2
			},
			src:             0,
			dst:             2,
			wantedDistances: []int64{0, 10, math.MaxInt64},
			wantedPrevious:  []int{-1, 0, -1},
		},
	}

	for _, c := range cases {
		distances, previous := myAlgorithms.DijkstraShortestPath(c.regionMap, c.src, c.dst)

		if len(c.wantedDistances) != len(distances) {
			t.Errorf(
				"len(DijkstraShortestPath(%v, %d, %d) => distances) == %d, want %d",
				c.regionMap,
				c.src,
				c.dst,
				len(distances),
				len(c.wantedDistances),
			)
		}

		if len(c.wantedPrevious) != len(previous) {
			t.Errorf(
				"len(DijkstraShortestPath(%v, %d, %d) => previous) == %d, want %d",
				c.regionMap,
				c.src,
				c.dst,
				len(previous),
				len(c.wantedPrevious),
			)
		}

		for i, distance := range distances {
			if distance != c.wantedDistances[i] {
				t.Errorf(
					"DijkstraShortestPath(%v, %d, %d) == %v, _, want %v",
					c.regionMap,
					c.src,
					c.dst,
					distances,
					c.wantedDistances,
				)
			}
		}

		for i, aPrevious := range previous {
			if aPrevious != c.wantedPrevious[i] {
				t.Errorf(
					"DijkstraShortestPath(%v, %d, %d) == _, %v, want %v",
					c.regionMap,
					c.src,
					c.dst,
					previous,
					c.wantedPrevious,
				)
			}
		}
	}
}
