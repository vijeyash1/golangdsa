package twosum

func Twosum(nums []int, target int) []int {
	length := len(nums)
	res := make([]int, 0)
	for i := 0; i < length; i++ {
		find := target - nums[i]
		for j := i + 1; j < length; j++ {

			if nums[j] == find {
				res = []int{i, j}
				return res
			}
		}
	}
	return res
}

func TwoSums(nums []int, target int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
