package Sqrt

func MySqrt(x int) int {
	left := 0
	right := x
	mid := 0
	for true {
		mid = (left + right) / 2
		if mid*mid == x {
			return mid
		} else {
			if mid*mid > x {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left > right && mid*mid > x {
			return mid - 1
		} else if left > right && mid*mid < x {
			return mid
		}
	}
	return mid
}
