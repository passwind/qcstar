package problems

import "strings"

func intToRoman(num int) string {
	t, i := num, 0
	rc := map[int]map[int]byte{
		0: map[int]byte{
			1: 'I',
			5: 'V',
		},
		1: map[int]byte{
			1: 'X',
			5: 'L',
		},
		2: map[int]byte{
			1: 'C',
			5: 'D',
		},
		3: map[int]byte{
			1: 'M',
		},
	}
	res := [][]byte{}
	for t > 0 {
		d := t % 10
		rr := []byte{}
		if d == 9 {
			rr = append(rr, rc[i][1])
			rr = append(rr, rc[i+1][1])
		} else if d == 4 {
			rr = append(rr, rc[i][1])
			rr = append(rr, rc[i][5])
		} else if d == 5 {
			rr = append(rr, rc[i][5])
		} else if d < 4 {
			for k := 0; k < d; k++ {
				rr = append(rr, rc[i][1])
			}
		} else {
			rr = append(rr, rc[i][5])
			for k := 6; k < 9; k++ {
				rr = append(rr, rc[i][1])
			}
		}
		res = append(res, rr)
		t = t / 10
		i++
	}
	str := []string{}
	for k := len(res) - 1; k >= 0; k-- {
		str = append(str, string(res[k]))
	}
	return strings.Join(str, "")
}
