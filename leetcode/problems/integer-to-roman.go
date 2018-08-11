package problems

func intToRoman(num int) string {
	t, i, j := num, 0, 127
	var res [128]byte
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
		} else if d < 4 && d > 0 {
			for k := 0; k < d; k++ {
				res[j] = rc[i][1]
				j--
			}
		} else if d >= 6 && d < 9 { // d>=6 && d<9
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
	return string(res[j+1:])
}
