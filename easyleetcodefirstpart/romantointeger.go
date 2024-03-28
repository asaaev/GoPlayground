package easyleetcodefirstpart

var romanLiterals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	var romanLiterals = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	totalSum := 0
	prevValue := 0
	for _, ch := range s {
		val := romanLiterals[ch]
		if val > prevValue {
			totalSum += val - 2*prevValue
		} else {
			totalSum += val
		}
		prevValue = val
	}
	return totalSum
}
