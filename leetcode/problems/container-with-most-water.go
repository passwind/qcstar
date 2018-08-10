package problems

func maxArea(height []int) int {
	total := len(height)
	var max, maxH int
	for i := total - 1; i >= 1; i-- {
		for j := 0; j+i < total; j++ {
			// fmt.Printf("i=%d, j=%d, %d - %d\n", i, j, height[j], height[j+i])
			if height[j] <= maxH || height[j+i] <= maxH {
				continue
			}
			cmh := height[j]
			if cmh > height[j+i] {
				cmh = height[j+i]
			}
			m := cmh * i
			if m > max {
				max = m
			}
			if cmh > maxH {
				maxH = cmh
			}
		}
	}

	return max
}

func maxArea2(height []int) int {
	total := len(height)
	var max int
	mh := make([]int, total-1)
	for i := 1; i < total; i++ {
		for j := 0; j+i < total; j++ {
			// fmt.Printf("i=%d, j=%d, %d - %d\n", i, j, height[j], height[j+i])
			cmh := height[j]
			if cmh > height[j+i] {
				cmh = height[j+i]
			}
			if mh[i-1] < cmh {
				mh[i-1] = cmh
			}
		}
	}
	// fmt.Printf("%+v\n", mh)
	for w, h := range mh {
		m := (w + 1) * h
		if m > max {
			max = m
		}
	}
	return max
}

func maxArea1(height []int) int {
	total := len(height)
	var max int
	for i := 0; i < total; i++ {
		for j := i + 1; j < total; j++ {
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			w := j - i
			m := w * h
			if m > max {
				max = m
			}
		}
	}
	return max
}
