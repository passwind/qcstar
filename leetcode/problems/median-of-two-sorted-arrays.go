package problems

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, l1, l2, x, y := 0, 0, len(nums1), len(nums2), 0, 0
	mid := (l1 + l2) / 2

	for {
		if i+j > mid {
			break
		}
		y = x
		if i < l1 && j < l2 {
			if nums1[i] >= nums2[j] {
				x = nums2[j]
				j++
			} else {
				x = nums1[i]
				i++
			}
		} else if i < l1 {
			x = nums1[i]
			i++
		} else {
			x = nums2[j]
			j++
		}
	}

	if (l1+l2)%2 != 0 {
		return float64(x)
	}

	return (float64(x + y)) / 2.0
}
