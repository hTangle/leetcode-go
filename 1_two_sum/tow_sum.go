package __two_sum

// twoSum https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	helper := map[int]int{}
	for index, num := range nums {
		t := target - num
		if val, ok := helper[t]; ok {
			return []int{index, val}
		}
		helper[num] = index
	}
	return []int{-1, -1}
}
