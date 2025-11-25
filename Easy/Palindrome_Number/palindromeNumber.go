package palindrome_number

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	first := x % 10
//	newNum := x
//	last := 0
//	counter := 0
//
//	for newNum > 0 {
//		last = newNum % 10
//		newNum = newNum / 10
//		counter++
//	}
//
//	newNum = (x - int(math.Pow10(counter-1))*last) / 10
//
//	if newNum < 10 && first == last {
//		return true
//	}
//
//	if first != last {
//		return false
//	}
//	return isPalindrome(newNum)
//}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	newNum := 0
	x1 := x
	for x1 > 0 {
		newNum = newNum*10 + x1%10
		x1 = x1 / 10
	}

	if x == newNum {
		return true
	}

	return false
}
