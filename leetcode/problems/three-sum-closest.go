package problems

import "math"

func threeSumClosest(nums []int, target int) int {
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
