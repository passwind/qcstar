package problems

var ric = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	str := []byte(s)
	var last byte
	res, ci := 0, 0

	last = '0'

	for _, c := range str {
		if c == last {
			ci += ric[c]
		} else if last != '0' {
			if ric[c] > ric[last] {
				ci = ric[c] - ci
				res += ci
				ci = 0
			} else {
				res += ci
				ci = ric[c]
			}
		} else {
			ci = ric[c]
		}
		last = c
	}
	res += ci

	return res
}
