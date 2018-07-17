package problems

func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}
	t := x
	if x < 0 {
		t = -x
	}
	var r int
	for t > 0 {
		tt := t % 10
		r = r*10 + tt
		t = (t - tt) / 10
	}
	if x < 0 {
		r = -r
	}

	if r < -2147483648 || r > 2147483647 {
		return 0
	}
	return r
}
