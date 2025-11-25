package MergeSortedArray

func Merge(nums1 []int, m int, nums2 []int, n int) {
	pointer1 := m - 1
	pointer2 := n - 1
	for i := len(nums1) - 1; i >= 0; i-- {
		if pointer2 < 0 {
			break
		}
		if pointer1 >= 0 && nums1[pointer1] >= nums2[pointer2] {
			nums1[i] = nums1[pointer1]
			pointer1--
		} else {
			nums1[i] = nums2[pointer2]
			pointer2--
		}
	}
}
