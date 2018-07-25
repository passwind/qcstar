package problems

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("s=%s, p=%s\n", s, p)
	totals, totalp := len(s), len(p)
	if totalp == 0 && totals > 0 {
		return false
	} else if totalp > 0 && totals == 0 {
		return false
	} else if totals == 0 && totalp == 0 {
		return true
	} else if s == p {
		return true
	}

	str := []byte(s)
	reg := []byte(p)
	i, j, start := 0, 0, 0
	for {
		fmt.Printf("start: b=%d, i=%d,r=%d, j=%d, start=%d\n", str[i], i, reg[j], j, start)
		if str[i] == reg[j] || reg[j] == '.' {
			start = j
			i++
			j++
		} else if reg[j] == '*' {
			if j == 0 {
				return false
			} else if str[i] == reg[start] || reg[start] == '.' {
				i++
			} else if str[i] != reg[start] && j < totalp-1 && (reg[j+1] == str[i] || reg[j+1] == '.') {
				j++
				start = j
			} else {
				return false
			}
		} else if str[i] != reg[j] && j < totalp-1 && reg[j+1] == '*' {
			j = j + 2
			start = j
		} else {
			return false
		}

		if i >= totals || j >= totalp {
			break
		}
	}

	fmt.Printf("check %d = %d, %d = %d\n", i, totals, j, totalp)

	if i == totals {
		if reg[totalp-1] == '*' {
			return j == totalp-1
		} else if j == totalp-2 && reg[j] == '*' && reg[start] == reg[totalp-1] {
			return true
		}
	}

	return i == totals && j == totalp
}
