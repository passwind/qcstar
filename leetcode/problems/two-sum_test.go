package problems

import "testing"

func Test_twoSumBasic(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	t.Logf("res is %+v", res)

	nums = []int{-3, 4, 3, 90}
	res = twoSum(nums, 0)
	t.Logf("res is %+v", res)
}
