package algorithms

import "math"

// SumNumbers sums a collection of numbers
func (a Algorithms) SumNumbers(numbers []float32) float32 {
	var sum float32

	for _, number := range numbers {
		sum += number
	}

	return (float32(math.Round(float64(sum)*100)) / 100)
}
