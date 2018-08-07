package problems

import "testing"

func Test_maxArea(t *testing.T) {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if res != 49 {
		t.Errorf("error: should be 49, in fact is %d", res)
	}
}
