package test

import (
	"goalgorithms"
	"testing"
)

func TestWordsCounter(t *testing.T) {
	myAlgorithms := goalgorithms.Algorithms{}
	cases := []struct {
		phrase string
		want   map[string]int
	}{
		{
			phrase: "",
			want:   make(map[string]int),
		},
		{
			phrase: "abc",
			want:   map[string]int{"abc": 1},
		},
		{
			phrase: "abc bcd efg",
			want:   map[string]int{"abc": 1, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd efg abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd, efg abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd efg. abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd, efg. abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc  bcd,, efg.... abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "    abc bcd efg abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc...bcd,,efg.,abc",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd efg abc abcd",
			want:   map[string]int{"abc": 2, "abcd": 1, "bcd": 1, "efg": 1},
		},
		{
			phrase: "efg abc bcd bcd bcd efg efg bcd abc",
			want:   map[string]int{"abc": 2, "bcd": 4, "efg": 3},
		},
		{
			phrase: "abc bcd efg abc ",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd efg abc,",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd efg abc.",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
		{
			phrase: "abc bcd efg abc ,.",
			want:   map[string]int{"abc": 2, "bcd": 1, "efg": 1},
		},
	}

	for _, c := range cases {
		counter := make(map[string]int)
		myAlgorithms.WordsCounter(c.phrase, counter)

		for k, v := range c.want {
			if counter[k] != v {
				t.Errorf(
					"WordsCounter(%v, {}) => counter[%v] == %d, want counter[%v] == %d",
					c.phrase,
					k,
					counter[k],
					k,
					v,
				)
			}
		}
	}
}
