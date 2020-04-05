package goalgorithms

// WordsCounter counts the number of occurrences for each word from the phrase
func (a Algorithms) WordsCounter(phrase string, counter map[string]int) {
	for i := 0; i < len(phrase); i++ {
		for ; i < len(phrase) && isDelimiter(phrase[i]); i++ {
		}

		if i < len(phrase) {
			first := i

			for ; i < len(phrase) && !isDelimiter(phrase[i]); i++ {
			}

			counter[phrase[first:i]]++
		}
	}
}

func isDelimiter(character byte) bool {
	delimiters := []byte{',', '.', ' '}

	for _, delimiter := range delimiters {
		if character == delimiter {
			return true
		}
	}

	return false
}
