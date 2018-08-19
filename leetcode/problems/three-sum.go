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
	i int
	j int
	k int
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	total := len(nums)
	if len(nums) < 3 {
		return res
	}

	quickSort(&nums, 0, len(nums)-1)
	// fmt.Printf("%+v\n", nums)

	fi, zi, zeroidx0, zeroidx1 := -1, total, -1, -1
	for k := total; k > 0; k-- {
		if nums[k-1] <= 0 {
			zi = k
			if nums[k-1] == 0 {
				zeroidx1 = k - 1
			}
			break
		}
	}
	for k := 0; k < total; k++ {
		if nums[k] >= 0 {
			fi = k - 1
			if nums[k] == 0 {
				zeroidx0 = k
			}
			break
		}
	}

	rm := make(map[threeSumKey]bool)

	fmt.Printf("%v - %d,%d,%d,%d - %d : %d\n", nums, fi, zeroidx0, zeroidx1, zi, total-1-zi, fi)

	if total-1 >= zi && fi >= 0 {
		for i := 1; i < total-1; i++ {
			if i == zeroidx0 {
				i = zi
			}
			for j := 0; j < i && j <= fi; j++ {
				if nums[j]+nums[i] < 0 {
					for k := len(nums) - 1; k >= i+1 && k >= zi; k-- {
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
						} else if nums[j]+nums[i]+nums[k] < 0 {
							break
						}
					}
				} else if zeroidx1 > 0 && nums[j]+nums[i] == 0 {
					key := threeSumKey{
						j: nums[j],
						i: 0,
						k: nums[i],
					}
					if _, ok := rm[key]; !ok {
						res = append(res, []int{nums[j], 0, nums[i]})
						rm[key] = true
					}
				} else {
					break
				}
			}
		}
	}

	if zeroidx1-zeroidx0 > 1 {
		res = append(res, []int{0, 0, 0})
	}

	return res
}
