package SearchInsertPosition

func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for true {
		mid = (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
		if left > right && nums[mid] > target {
			return mid
		} else if left > right && nums[mid] < target {
			return mid + 1
		}

	}
	return 0
}
