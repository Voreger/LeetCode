package AddBinary

import "slices"

func AddBinary(a string, b string) string {
	remains := 0
	A := []rune(a)
	B := []rune(b)
	slices.Reverse(A)
	slices.Reverse(B)
	res := []rune{}
	for i := 0; i < len(A) || i < len(B); i++ {
		sum := 0
		if i < len(A) && A[i] == '1' {
			sum++
		}
		if i < len(B) && B[i] == '1' {
			sum++
		}
		sum += remains
		if sum == 3 {
			res = append(res, '1')
			remains = 1
		} else if sum == 2 {
			res = append(res, '0')
			remains = 1
		} else if sum == 1 {
			res = append(res, '1')
			remains = 0
		} else {
			res = append(res, '0')
			remains = 0
		}
	}
	if remains > 0 {
		res = append(res, '1')
	}
	slices.Reverse(res)
	r := string(res)
	return r

}
