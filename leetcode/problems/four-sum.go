package problems

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	r := make([][]int, 0)
	if l < 4 {
		return r
	}
	// fmt.Printf("in %v\n", nums)
	m := nums[0] - 1
	for i := 0; i < l-3; i++ {
		if m == nums[i] {
			continue
		}
		m = nums[i]
		t := target - nums[i]
		n := nums[i+1] - 1
		for j := i + 1; j < l-2; j++ {
			if n == nums[j] {
				continue
			}
			n = nums[j]
			tt := t - nums[j]
			x := j + 1
			y := l - 1
			for x < y {
				ttt := nums[x] + nums[y]
				// fmt.Printf("[i=%d, j=%d, x=%d, y=%d],target=%d, t=%d, tt=%d, ttt=%d\n", i, j, x, y, target, t, tt, ttt)
				if ttt == tt {
					r = append(r, []int{nums[i], nums[j], nums[x], nums[y]})
					x0 := x
					y0 := y
					for x < l-1 {
						x++
						if nums[x0] != nums[x] {
							break
						}
					}
					for y > j {
						y--
						if nums[y0] != nums[y] {
							break
						}
					}
				} else if ttt > tt {
					y--
				} else if ttt < tt {
					x++
				}
			}
		}
	}
	return r
}
