package test

import (
	"algorithms"
	"testing"
)

func TestMultiplyNumbers(t *testing.T) {
	var (
		myAlgorithms = algorithms.Algorithms{}
		cases        = []struct {
			numbers []float32
			want    float32
		}{
			{numbers: []float32{1.0, 2.0, 3.0}, want: 6.0},
			{numbers: []float32{1.2, 3.4, 8.8, 9.56}, want: 343.24},
			{numbers: []float32{}, want: 0.0},
		}
	)

	for _, c := range cases {
		got := myAlgorithms.MultiplyNumbers(c.numbers)

		if got != c.want {
			t.Errorf("MultiplyNumbers(%f) == %f, want %f", c.numbers, got, c.want)
		}
	}
}
