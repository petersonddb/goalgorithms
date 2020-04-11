package goalgorithms

type positionStatus int

const (
	out positionStatus = iota
	in
)

// WordsCounter counts the number of occurrences for each word from the phrase
func (a Algorithms) WordsCounter(phrase string, counter map[string]int) {
	var first int
	status := out
	n := len(phrase)

	for i := 0; i <= n; i++ {
		if i == n || isDelimiter(phrase[i]) {
			if status == in {
				status = out
				counter[phrase[first:i]]++
			}
		} else if status == out {
			status = in
			first = i
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
