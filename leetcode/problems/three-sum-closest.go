package problems

import (
	"math"
	"sort"
)

func abs(v int) int {
	if v < 0 {
		return 0 - v
	}
	return v
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)

	if l < 3 {
		return -1
	}

	r := nums[0] + nums[1] + nums[2]

	m := nums[0] - 1
	for i := 0; i < l-2; i++ {
		if m == nums[i] {
			continue
		}
		m = nums[i]
		tt := target - nums[i]

		j := i + 1
		k := l - 1
		for j < k {
			tr := m + nums[j] + nums[k]
			if abs(tr-target) < abs(r-target) {
				r = tr
			}
			if nums[j]+nums[k] > tt {
				k--
			} else if nums[j]+nums[k] < tt {
				j++
			} else {
				return target
			}
		}
	}

	return r
}

func threeSumClosest1(nums []int, target int) int {
	min := math.MaxInt32
	res := math.MinInt32
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				t := nums[i] + nums[j] + nums[k]
				if math.Abs(float64(t-target)) < float64(min) {
					min = int(math.Abs(float64(t - target)))
					res = t
				}
			}
		}
	}
	return res
}
