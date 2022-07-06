package solutions

func TwoSum(nums []int, target int) []int {
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if m[nums[i]] != 0 {
			return []int{m[nums[i]] - 1, i}
		} else {
			m[sub] = i + 1
		}
	}
	return []int{}
}
