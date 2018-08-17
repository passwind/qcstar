package problems

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	i, stop := 0, false
	for {
		for _, s := range strs[1:] {
			if i >= len(strs[0]) || i >= len(s) || strs[0][i] != s[i] {
				stop = true
				break
			}
		}
		if stop {
			break
		}
		i++
	}

	return strs[0][:i]
}
