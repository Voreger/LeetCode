package AddBinary

import "slices"

func AddBinary(a string, b string) string {
	remains := 0
	aRune := []rune(a)
	bRune := []rune(b)
	slices.Reverse(aRune)
	slices.Reverse(bRune)
	result := []rune{}
	for i := 0; i < len(aRune) || i < len(bRune); i++ {
		sum := 0
		if i < len(aRune) && aRune[i] == '1' {
			sum++
		}
		if i < len(bRune) && bRune[i] == '1' {
			sum++
		}
		sum += remains
		result = append(result, rune(sum%2)+'0')
		remains = sum / 2
	}
	if remains == 1 {
		result = append(result, '1')
	}
	slices.Reverse(result)
	return string(result)
}
