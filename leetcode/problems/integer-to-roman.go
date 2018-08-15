package problems

import (
	"bytes"
)

var rc = map[int]map[int]byte{
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

func intToRoman(num int) string {
	t, i, j := num, 3, 1000
	b := bytes.NewBuffer(nil)

	for t > 0 {
		d := t / j
		t = t - d*j
		j = j / 10
		if d == 9 {
			b.WriteByte(rc[i][1])
			b.WriteByte(rc[i+1][1])
		} else if d == 4 {
			b.WriteByte(rc[i][1])
			b.WriteByte(rc[i][5])
		} else if d == 5 {
			b.WriteByte(rc[i][5])
		} else if d >= 1 && d <= 3 {
			for k := 0; k < d; k++ {
				b.WriteByte(rc[i][1])
			}
		} else if d >= 6 && d <= 8 {
			b.WriteByte(rc[i][5])
			for k := 6; k <= d; k++ {
				b.WriteByte(rc[i][1])
			}
		}
		i--
	}
	return b.String()
}

func intToRoman1(num int) string {
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
