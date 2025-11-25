package ClimbingStairs

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	firstStair, secondStair := 1, 2
	for i := 3; i <= n; i++ {
		firstStair, secondStair = secondStair, firstStair+secondStair
	}
	return secondStair
}
