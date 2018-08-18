package problems

import "fmt"

func quickSort(in *[]int, begin, end int) {
	var partition = func(b, e int) int {
		i, j, pivot := b, e, b
		pivotnumber := (*in)[pivot]
		for i != j {
			var tmp int
			for (*in)[j] > pivotnumber && i < j {
				j--
			}
			for (*in)[i] <= pivotnumber && i < j {
				i++
			}
			tmp = (*in)[i]
			(*in)[i] = (*in)[j]
			(*in)[j] = tmp
		}
		if i == j {
			tmp := (*in)[pivot]
			(*in)[pivot] = (*in)[i]
			(*in)[i] = tmp
		}
		return i
	}

	if begin < end {
		m := partition(begin, end)
		quickSort(in, begin, m-1)
		quickSort(in, m+1, end)
	}
}

type threeSumKey struct {
	j int
	i int
	k int
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	quickSort(&nums, 0, len(nums)-1)
	fmt.Printf("%+v\n", nums)

	rm := make(map[threeSumKey]bool)
	for i := 1; i < len(nums)-1; i++ {
		for j := 0; j < i; j++ {
			if nums[j]+nums[i] <= 0 {
				for k := i + 1; k < len(nums); k++ {
					if nums[j]+nums[i]+nums[k] == 0 {
						key := threeSumKey{
							j: nums[j],
							i: nums[i],
							k: nums[k],
						}
						if _, ok := rm[key]; !ok {
							res = append(res, []int{nums[j], nums[i], nums[k]})
							rm[key] = true
						}
					} else if nums[j]+nums[i]+nums[k] > 0 {
						break
					}
				}
			} else {
				break
			}
		}
	}
	return res
}
