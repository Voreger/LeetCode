package TwoSum

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}
