package problems

import (
	"math"
)

func myAtoi(str string) int {
	bs := []byte(str)
	bflag := true
	start := false
	var res int
	for _, b := range bs {
		if !start && b == ' ' {
			continue
		} else if !start && (b == '+' || b == '-') {
			start = true
			if b == '-' {
				bflag = false
			}
			continue
		} else if !start && b >= '0' && b <= '9' {
			start = true
		} else if b < '0' || b > '9' {
			break
		}
		res = res*10 + int(b-'0')
		tmp := res
		if !bflag && tmp > 0 {
			tmp = -tmp
		}
		if tmp < math.MinInt32 {
			res = math.MinInt32
			break
		}
		if tmp > math.MaxInt32 {
			res = math.MaxInt32
			break
		}
	}
	if !bflag && res > 0 {
		return -res
	}
	return res
}
