package problems

import "fmt"

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

func longestPalindrome1(s string) string {
	total := len(s)
	if len(s) <= 1 {
		return s
	}

	idx := 0
	maxLen := 1
	start := 0
	for {
		if maxLen/2 >= total-idx {
			break
		}
		l, r := idx, idx
		for r+1 < total && s[r] == s[r+1] {
			r++
		}
		idx = r + 1
		for l > 0 && r < total-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if maxLen < r-l+1 {
			maxLen = r - l + 1
			start = l
		}
	}
	return s[start : start+maxLen]
}

func longestPalindromebb(s string) string {
	str := []byte(s)
	if len(s) < 2 {
		return s
	}
	index := 0
	maxLen, start := 1, 0
	for index < len(s) {
		if maxLen/2 >= len(s)-index {
			break
		}
		j, k := index, index
		for k < len(s)-1 && str[k+1] == str[k] {
			k++
		}
		index = k + 1
		for k < len(s)-1 && j > 0 && str[j-1] == str[k+1] {
			j--
			k++
		}
		fmt.Printf("hh %s %d %d %d %d\n", s, j, k, maxLen, index)
		if k-j+1 > maxLen {
			start = j
			maxLen = k - j + 1
		}
	}
	return (string(str[start : start+maxLen]))
}
