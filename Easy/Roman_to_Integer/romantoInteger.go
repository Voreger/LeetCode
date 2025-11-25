package Roman_to_Integer

import (
	"strings"
)

func RomanToInt(s string) int {
	var sum int

	rim := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	rimLow := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	for i, v := range rimLow {
		sum += strings.Count(s, i) * v
		s = strings.ReplaceAll(s, i, "")
	}

	for i, v := range rim {
		sum += strings.Count(s, i) * v
	}

	return sum
}
