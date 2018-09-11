package problems

import "testing"

func Test_fourSum(t *testing.T) {
	// nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
	// 	满足要求的四元组集合为：
	// [
	//   [-1,  0, 0, 1],
	//   [-2, -1, 1, 2],
	//   [-2,  0, 0, 2]
	// ]
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	t.Logf("res is %v", res)
}
