package problems

import "bytes"

func intToRoman(num int) string {
	t, i, j := num, 0, 127
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
	res := [128]byte{}
	for t > 0 {
		d := t % 10
		if d == 9 {
			res[j] = rc[i+1][1]
			j--
			res[j] = rc[i][1]
			j--
		} else if d == 4 {
			res[j] = rc[i][5]
			j--
			res[j] = rc[i][1]
			j--
		} else if d == 5 {
			res[j] = rc[i][5]
			j--
		} else if d >= 1 && d <= 3 {
			for k := 0; k < d; k++ {
				res[j] = rc[i][1]
				j--
			}
		} else if d >= 6 && d <= 8 {
			for k := 6; k <= d; k++ {
				res[j] = rc[i][1]
				j--
			}
			res[j] = rc[i][5]
			j--
		}
		t = t / 10
		i++
	}
	b := bytes.NewBuffer(res[j+1:])
	return b.String()
}
