package SmallestIntegerDivisibleByK

func SmallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%10 == 5 {
		return -1
	}
	i := 1
	length := 1
	for i%k != 0 {
		i = (i%k)*10 + 1
		length++
	}
	return length
}
