package problems

import "testing"

func Test_threeSumClosest(t *testing.T) {
	res := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	if res != 2 {
		t.Errorf("error: [-1，2，1，-4] should be 2, in fact [%d]", res)
	}
}
