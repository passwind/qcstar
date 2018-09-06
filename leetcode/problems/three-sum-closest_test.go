package problems

import "testing"

func Test_threeSumClosest(t *testing.T) {
	res := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	if res != 2 {
		t.Errorf("error: [-1，2，1，-4] should be 2, in fact [%d]", res)
	}

	// 新算法
	// 1.[1,1,-1,-1,3]
	// 3
	// 输出：
	// 1
	// 预期：
	// 3
	res = threeSumClosest([]int{1, 1, -1, -1, 3}, 3)
	if res != 3 {
		t.Errorf("error: [1,1,-1,-1,3] should be 3, in fact [%d]", res)
	}
}
