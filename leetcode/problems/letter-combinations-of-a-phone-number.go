package problems

import (
	"bytes"
)

var codeMap = map[byte][]byte{
	'2': []byte{
		'a',
		'b',
		'c',
	},
	'3': []byte{
		'd',
		'e',
		'f',
	},
	'4': []byte{
		'g',
		'h',
		'i',
	},
	'5': []byte{
		'j',
		'k',
		'l',
	},
	'6': []byte{
		'm',
		'n',
		'o',
	},
	'7': []byte{
		'p',
		'q',
		'r',
		's',
	},
	'8': []byte{
		't',
		'u',
		'v',
	},
	'9': []byte{
		'w',
		'x',
		'y',
		'z',
	},
}

func letterCombinations(digits string) []string {
	str := []byte(digits)
	arr := []*bytes.Buffer{}
	for i, b := range str {
		if i == 0 {
			for _, c := range codeMap[b] {
				a := bytes.NewBuffer(nil)
				a.WriteByte(c)
				arr = append(arr, a)
			}
		} else {
			arr1 := []*bytes.Buffer{}
			for _, aa := range arr {
				for _, c := range codeMap[b] {
					aaa := bytes.NewBuffer([]byte(aa.String()))
					aaa.WriteByte(c)
					arr1 = append(arr1, aaa)
				}
			}
			arr = arr1
		}
	}
	res := []string{}
	for _, ar := range arr {
		res = append(res, ar.String())
	}
	return res
}
