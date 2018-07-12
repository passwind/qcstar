package problems

import "fmt"

func longestPalindrome(s string) string {
	fmt.Printf("string is %s\n", s)
	if len(s) <= 1 {
		return s
	}

	l, r, m, ml, mr := 0, 0, 1, 0, 0
	var t byte
	for idx, c := range s {
		fmt.Printf("l, r, m, ml, mr, idx = %d, %d, %d, %d, %d, %d\n", l, r, m, ml, mr, idx)
		b := byte(c)
		if b == t {
			r = idx
			m++
		} else {
			t = b
			if idx-(2*m-1) >= 0 && b == s[idx-(2*m-1)] { // 无中点对称, 如bbbb，aa
				l = idx - (2*m - 1)
				r = idx
				m++
			} else if idx-2*m >= 0 && b == s[idx-2*m] { // 有中点对称，如bbb，aba
				l = idx - 2*m
				r = idx
				m++
			} else {
				m = 1
				l = idx
			}
		}

		if mr-ml < r-l {
			ml = l
			mr = r
		}
	}

	fmt.Printf("ml, mr = %d, %d\n", ml, mr)

	return s[ml : mr+1]
}
