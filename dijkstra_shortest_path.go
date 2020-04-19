package goalgorithms

import "math"

const (
	noPath int64 = -1
	noOne  int   = -1
)

// DijkstraShortestPath calculates the shortest path
// from a src to every destination in a region map
func (a Algorithms) DijkstraShortestPath(regionMap [][]int64, src int, dst int) ([]int64, []int) {
	n := len(regionMap)
	distances := make([]int64, n)
	previous := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		distances[i] = math.MaxInt64
		previous[i] = noOne
	}
	distances[src] = 0

	for target := nextMinNotVisited(distances, visited); target != noOne; target = nextMinNotVisited(distances, visited) {
		visited[target] = true

		if target == dst {
			break
		}

		for adj := 0; adj < len(regionMap[target]); adj++ {
			if regionMap[target][adj] == noPath {
				continue
			}

			distance := distances[target] + regionMap[target][adj]
			if distance < distances[adj] {
				distances[adj] = distance
				previous[adj] = target
			}
		}
	}

	return distances, previous
}

func nextMinNotVisited(distances []int64, visited []bool) int {
	minNotVisited, minDistance := noOne, int64(math.MaxInt64)

	for i, distance := range distances {
		if !visited[i] && distance < minDistance {
			minNotVisited, minDistance = i, distance
		}
	}

	return minNotVisited
}
