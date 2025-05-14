package RemoveElement

func RemoveElement(nums []int, val int) int {
	ind := -1

	for _, num := range nums {
		if num != val {
			ind++
			nums[ind] = num
		}
	}

	return ind + 1
}
