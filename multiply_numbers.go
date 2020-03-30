package goalgorithms

import "math"

// MultiplyNumbers multiplies a collection of numbers
func (a Algorithms) MultiplyNumbers(numbers []float32) float32 {
	var multiply float32

	if len(numbers) > 0 {
		multiply = 1.0
		for _, number := range numbers {
			multiply *= number
		}
	}

	return (float32(math.Round(float64(multiply)*100)) / 100)
}
