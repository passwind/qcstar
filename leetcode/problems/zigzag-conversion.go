package problems

import (
	"strings"
)

func convert(s string, numRows int) string {
	// fmt.Printf("s=%s,row=%d\n", s, numRows)
	ss := [][]byte{}
	for i := 0; i < numRows; i++ {
		ss = append(ss, []byte{})
	}
	row := 0
	downflag := true
	for _, b := range []byte(s) {
		ss[row] = append(ss[row], b)

		if downflag {
			row++
		} else {
			row--
		}

		if row == numRows {
			downflag = false
			row = row - 2
			if row < 0 {
				row = 0
			}
		}

		if row == 0 {
			downflag = true
		}
		// fmt.Printf("downflag=%t, row=%d\n", downflag, row)
	}
	strs := []string{}
	for _, sss := range ss {
		strs = append(strs, string(sss))
	}
	return strings.Join(strs, "")
}
