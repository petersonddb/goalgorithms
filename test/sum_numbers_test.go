package test

import (
	"algorithms"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	var (
		myAlgorithms = algorithms.Algorithms{}
		cases        = []struct {
			numbers []float32
			want    float32
		}{
			{numbers: []float32{1.0, 2.0, 3.0}, want: 6.0},
			{numbers: []float32{1.2, 3.4, 8.8, 9.56}, want: 22.96},
			{numbers: []float32{}, want: 0.0},
		}
	)

	for _, c := range cases {
		got := myAlgorithms.SumNumbers(c.numbers)

		if got != c.want {
			t.Errorf("SumNumbers(%f) == %f, want %f", c.numbers, got, c.want)
		}
	}
}
