package problems

func checkmatch0(str, reg []byte, si0, si1, sj0, sj1, ts, tp int) bool {
	// fmt.Printf("check1111: %d, %d - %d, %d\n", si0, si1, sj0, sj1)
	if (si0 == ts && sj0 == tp) || (si1 < si0 && sj1 < sj0) {
		return true
	}
	if si1 >= si0 && sj0 > sj1 {
		return false
	}
	x, y := si0, sj0
	for y < sj1 && reg[y+1] == '*' {
		for x <= si1 && (reg[y] == '.' || reg[y] == str[x]) {
			x++
		}
		y = y + 2
	}
	// fmt.Printf("c0: x=%d, y=%d\n", x, y)
	return x == si1+1 && y == sj1+1
}

func checkmatch(str, reg []byte, si0, si1, sj0, sj1, ts, tp int) bool {
	// fmt.Printf("check: %d, %d - %d, %d\n", si0, si1, sj0, sj1)
	if (si0 == ts && sj0 == tp) || (si1 < si0 && sj1 < sj0) {
		return true
	}
	if si1 >= si0 && sj0 > sj1 {
		return false
	}
	start, end := -1, -1
	for k := sj0; k <= sj1; k++ {
		if k < sj1 {
			if start > -1 && reg[k+1] == '*' {
				end = k - 1
				break
			} else if reg[k+1] != '*' && start == -1 {
				start = k
				end = k
			} else if reg[k+1] == '*' {
				k++
			}
		} else if k == sj1 && start == -1 {
			start = k
			end = k
		}
	}

	if start == -1 {
		return checkmatch0(str, reg, si0, si1, sj0, sj1, ts, tp)
	}

	if si1-si0 < end-start {
		return false
	}

	sii := si0
	for sii <= si1 {
		if reg[start] == '.' || reg[start] == str[sii] {
			sfind := true
			for k := start; k <= end; k++ {
				if sii+(k-start) == ts || (reg[k] != '.' && reg[k] != str[sii+(k-start)]) {
					sfind = false
					break
				}
			}
			if !sfind {
				sii++
				continue
			}
		} else {
			sii++
			continue
		}

		s1find := checkmatch0(str, reg, si0, sii-1, sj0, start-1, ts, tp)
		s2find := checkmatch(str, reg, sii+(end-start+1), si1, end+1, sj1, ts, tp)

		// fmt.Printf("1=%t, 2=%t\n", s1find, s2find)

		if s1find && s2find {
			return true
		}
		sii++
	}

	return false
}

func isMatch(s string, p string) bool {
	// fmt.Printf("s=%s, p=%s\n", s, p)
	totals, totalp := len(s), len(p)
	str := []byte(s)
	reg := []byte(p)

	return checkmatch(str, reg, 0, totals-1, 0, totalp-1, totals, totalp)
}
