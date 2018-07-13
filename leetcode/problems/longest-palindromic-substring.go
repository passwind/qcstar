package problems

func longestPalindrome(s string) string {
	total := len(s)
	if len(s) <= 1 {
		return s
	}

	checkMax := func(ml int, idx int) bool {
		l := idx - ml/2
		r := idx + ml/2
		if ml%2 == 0 {
			r = idx + ml/2 - 1
		}
		i, j := l, r
		if l < 0 || r >= total {
			return false
		}
		for {
			if i >= j {
				break
			}
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}

		return true
	}

	maxlen := 1
	maxstr := s[0:1]

	for idx := 1; idx < total; idx++ {
		substr, ok := func() (string, bool) {
			ml1, ml2 := maxlen+1, maxlen+2
			ml1ok := checkMax(ml1, idx)
			ml2ok := checkMax(ml2, idx)
			var ml int
			if ml1ok && ml2ok {
				if ml1 > ml2 {
					ml = ml1
				} else {
					ml = ml2
				}
			} else if ml1ok {
				ml = ml1
			} else if ml2ok {
				ml = ml2
			} else {
				return "", false
			}
			var l, r int
			if ml%2 == 0 {
				l = idx - ml/2
				r = idx + ml/2 - 1
			} else {
				l = idx - ml/2
				r = idx + ml/2
			}

			i := l - 1
			j := r + 1
			for {
				if i < 0 || j == total || s[i] != s[j] {
					break
				}
				i--
				j++
			}
			return s[i+1 : j], true
		}()
		if ok {
			maxlen = len(substr)
			maxstr = substr
		}
	}

	return maxstr
}
