package problems

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

func twoSum1(nums []int, target int) [][]int {
	result := [][]int{}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			res := make([]int, 2)
			res[1] = i
			res[0] = m[target-nums[i]]
			result = append(result, res)
		}
		m[nums[i]] = i
	}
	return result
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	total := len(nums)
	if len(nums) < 3 {
		return res
	}

	quickSort(&nums, 0, len(nums)-1)
	// fmt.Printf("%+v\n", nums)

	fb, fi, zi, ze, zeroidx0, zeroidx1 := -1, -1, total, total, -1, -1
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
	tm := make(map[int]int)

	// fmt.Printf("%v - fb=%d, fi=%d, zeroidx0=%d, zeroidx1=%d, zi=%d, ze=%d - %d : %d\n", nums, fb, fi, zeroidx0, zeroidx1, zi, ze, total-1-zi, fi)

	if total-1 >= zi && fi >= 0 {
		for i := 0; i < total; i++ {
			if i == zeroidx0 {
				i = zi
			}
			if i <= fi {
				target := -nums[i]
				// fmt.Printf("fff: %v, i=%d, target=%d\n", nums[zi:], i, target)
				ze = total
				for nums[ze-1] >= target && ze > 0 {
					ze--
				}
				matched := twoSum1(nums[zi:ze], target)
				if len(matched) > 0 {
					for _, zis := range matched {
						key := threeSumKey{
							j: nums[i],
							i: nums[zi+zis[0]],
							k: nums[zi+zis[1]],
						}
						if _, ok := rm[key]; !ok {
							res = append(res, []int{nums[i], nums[zi+zis[0]], nums[zi+zis[1]]})
							rm[key] = true
						}
					}
				}
				if zeroidx1 > 0 {
					tm[-nums[i]] = i
				}
				// fmt.Printf("zis:=%d , res: %v\n", zis, res)
			} else if i >= zi {
				target := -nums[i]
				// fmt.Printf("zzz: %v, i=%d, target=%d\n", nums[0:fi+1], i, target)
				fb = 0
				for nums[fb] <= target && fb < total-1 {
					fb++
				}
				// fmt.Printf("fb=%d, fi=%d\n", fb, fi)
				matched := twoSum1(nums[fb:fi+1], target)
				if len(matched) > 0 {
					for _, zis := range matched {
						key := threeSumKey{
							j: nums[fb+zis[0]],
							i: nums[fb+zis[1]],
							k: nums[i],
						}
						if _, ok := rm[key]; !ok {
							res = append(res, []int{nums[fb+zis[0]], nums[fb+zis[1]], nums[i]})
							rm[key] = true
						}
					}
					// fmt.Printf("res: %v\n", res)
				}
				if zeroidx1 > 0 {
					if _, ok := tm[nums[i]]; ok {
						key := threeSumKey{
							j: nums[tm[nums[i]]],
							i: 0,
							k: nums[i],
						}
						if _, ok := rm[key]; !ok {
							res = append(res, []int{nums[tm[nums[i]]], 0, nums[i]})
							rm[key] = true
						}
					}
				}
			}
		}
	}

	if zeroidx1-zeroidx0 > 1 {
		res = append(res, []int{0, 0, 0})
	}

	return res
}
