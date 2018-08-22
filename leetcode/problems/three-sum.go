package problems

type threeSumKey struct {
	i int
	j int
	k int
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	zint, fint, intmap := []int{}, []int{}, map[int]int{}
	rm := make(map[threeSumKey]bool)
	zeronums := 0
	for idx, n := range nums {
		if n > 0 {
			zint = append(zint, idx)
			intmap[n] = idx
		} else if n < 0 {
			fint = append(fint, idx)
			intmap[n] = idx
		} else {
			zeronums++
		}
	}

	// fmt.Printf("%v - z: %v, f: %v, %d\n", nums, zint, fint, zeronums)

	for i := 0; i < len(fint); i++ {
		fn := nums[fint[i]]
		for j := 0; j < len(zint); j++ {
			zn := nums[zint[j]]
			var n1, n2, n3 int
			find := false
			d := -(fn + zn)
			if d == 0 && zeronums > 0 {
				n1 = fn
				n2 = 0
				n3 = zn
				find = true
			} else if idx, ok := intmap[d]; ok && idx != fint[i] && idx != zint[j] {
				if d < fn {
					n1 = d
					n2 = fn
					n3 = zn
				} else if d > zn {
					n1 = fn
					n2 = zn
					n3 = d
				} else {
					n1 = fn
					n2 = d
					n3 = zn
				}
				find = true
			}
			if find {
				key := threeSumKey{
					j: n1,
					i: n2,
					k: n3,
				}
				if _, ok := rm[key]; !ok {
					res = append(res, []int{n1, n2, n3})
					rm[key] = true
				}
			}
		}
	}
	if zeronums >= 3 {
		res = append(res, []int{0, 0, 0})
	}
	return res
}
