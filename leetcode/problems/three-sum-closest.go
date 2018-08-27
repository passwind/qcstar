package problems

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	fmt.Printf("haha: %v - %d\n", nums, target)
	for i := 0; i < len(nums)-2; i++ {
		nums[i] = nums[i] + nums[i+1] + nums[i+2]
	}
	nn := nums[0 : len(nums)-2]
	fmt.Printf("before: %v\n", nn)
	sort.Ints(nn)
	fmt.Printf("after: %v\n", nn)
	var i int
	for i = 0; i < len(nn); i++ {
		if nn[i]-target == 0 {
			return nn[i]
		} else if nn[i]-target > 0 {
			if i > 0 && target-nn[i-1] < nn[i]-target {
				return nn[i-1]
			}
			return nn[i]
		}
	}

	return nn[len(nn)-1]
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
