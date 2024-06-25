package twosum

import "fmt"

// For single array
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

// for single array
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

func TwoSumForTwoArray(arr1, arr2 []int, target int) {
	m := len(arr1)
	n := len(arr2)
	found := false

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr1[i]+arr2[j] == target {
				found = true
				fmt.Printf("the array 1 has index value of %v at %v and array 2 has index value of %v at %v and both adds up to target value %v", i, arr1[i], j, arr2[j], arr1[i]+arr2[j])
				return
			}
		}
	}

	if !found {
		fmt.Println("the sum is not found")
	}
}
