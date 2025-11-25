package PlusOne

//func PlusOne(digits []int) []int {
//	var result []int
//	var num int
//	count := 0
//	slices.Reverse(digits)
//	for _, digit := range digits {
//		num += digit * int(math.Pow(10, float64(count)))
//		count++
//	}
//	num = num + 1
//	for num > 0 {
//		result = append(result, num%10)
//		num /= 10
//	}
//	slices.Reverse(result)
//	return result
//}

func PlusOne(digits []int) []int {
	additional := 1

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+additional > 9 {
			digits[i] = 0
			additional = 1
		} else {
			digits[i] += additional
			additional = 0
		}
	}

	if additional == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
