package problems

import "testing"

func Test_maxArea(t *testing.T) {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if res != 49 {
		t.Errorf("error: should be 49, in fact is %d", res)
	}

	// 第二种方法
	// 1.
	res = maxArea([]int{2, 3, 4, 5, 18, 17, 6})
	if res != 17 {
		t.Errorf("error: should be 17, in fact is %d", res)
	}

}
