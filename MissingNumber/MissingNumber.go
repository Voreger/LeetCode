package MissingNumber

func MissingNumber(nums []int) int {
	n := len(nums)
	sum := nums[n-1] / 2 * (n + 1)
	return sum
}
