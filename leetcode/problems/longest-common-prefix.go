package problems

import (
	"bytes"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	res := bytes.NewBuffer(nil)

	i, stop := 0, false
	for len(strs) > 1 {
		for _, s := range strs[1:] {
			// fmt.Printf("i=%d, 0=%s, s=%s\n", i, strs[0], s)
			if i >= len(strs[0]) || i >= len(s) {
				stop = true
				break
			}
			if strs[0][i] != s[i] {
				stop = true
				break
			}
		}
		if stop {
			break
		}
		res.WriteByte(strs[0][i])
		i++
	}

	return res.String()
}
