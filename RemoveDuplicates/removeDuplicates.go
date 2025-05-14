package RemoveDuplicates

func RemoveDuplicates(nums []int) int {
	count := 0
	for i, num := range nums {
		if i != 0 && num == nums[i-1] {
			count++
		}
		if i != 0 {
			nums[i-count] = num
		}

	}
	return len(nums) - count
}
