package problems

func maxArea(height []int) int {
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
