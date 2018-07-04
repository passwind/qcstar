package problems

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 1
	for i := 0; i < len(s)-maxLen; i++ {
		tmpLen := 1
		for {
			nextIdx := i + tmpLen
			if nextIdx == len(s) {
				break
			}

			tmpStr := s[i:nextIdx]
			var nextChar string
			if nextIdx+1 == len(s) {
				nextChar = s[nextIdx:]
			} else {
				nextChar = s[nextIdx : nextIdx+1]
			}

			if strings.Index(tmpStr, nextChar) >= 0 {
				break
			}
			tmpLen++
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
		}
	}

	return maxLen
}

func nblengthOfLongestSubstring(s string) int {
	last, max, count := 1, 0, 0
	m := [128]int{}
	for index, value := range s {
		if m[value] < last {
			count += 1
		} else {
			last = m[value]
			if count > max {
				max = count
			}
			count = index - last + 1
		}
		m[value] = index + 1
	}
	if count > max {
		max = count
	}
	return max
}
