package problems

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("s=%s, p=%s\n", s, p)
	totals, totalp := len(s), len(p)
	str := []byte(s)
	reg := []byte(p)

	if totalp == 0 && totals > 0 {
		return false
	} else if totalp%2 == 1 && totals == 0 {
		return false
	} else if totalp%2 == 0 && totals == 0 {
		// "", "a*"
		// "", ".*"
		// "", "a*.*"
		for k := 0; k < totalp-1; k += 2 {
			if reg[k+1] != '*' {
				return false
			}
		}
		return true
	} else if totals == 0 && totalp == 0 {
		return true
	} else if s == p {
		return true
	}

	regstr := make([]int, totalp)
	i, j := 0, 0

	var backcheck = func() bool {
		fmt.Printf("backcheck: i=%d, j=%d, %v\n", i, j, regstr)
		if j < totalp-1 && reg[j+1] == '*' {
			j = j + 2
			return true
		}
		l := j
		find := false
		for k := j + 1; k < totalp; k++ {
			if k < totalp-1 && reg[k+1] == '*' {
				l = k - 1
				break
			} else if reg[j] == reg[k] {
				l = k
			} else {
				break
			}
		}
		fmt.Printf("l=%d,j=%d\n", l, j)
		len := l - j + 1
		if j > 0 {
			for k := j - 1; k > 0; k-- {
				fmt.Printf("backstr: k=%d, i=%d, j=%d\n", k, i, j)
				if reg[k] == '*' && (reg[k-1] == reg[j] || reg[k-1] == '.') {
					if regstr[k-1] > -1 {
						len1 := 0
						if reg[k-1] == '.' && reg[j] != '.' {
							end := -1
							for x := regstr[k]; x >= regstr[k-1] && len1 < len; x-- {
								fmt.Printf("len1=%d, x=%d, end=%d\n", len1, x, end)
								if str[x] == reg[j] {
									len1++
									if end == -1 {
										end = x
									}
								} else {
									len1 = 0
									end = -1
								}
							}
							if len > len1 {
								return false
							}
							i = end - len1 + 1
							if i == regstr[k-1] {
								regstr[k] = -1
								regstr[k-1] = -1
							} else {
								regstr[k] = i - 1
							}
						} else {
							len1 = regstr[k] - regstr[k-1] + 1
							if len > len1 {
								return false
							}
							i = regstr[k] - len + 1
							fmt.Printf("k=%d, i=%d, j=%d\n", k, i, j)
							if len == len1 {
								regstr[k] = -1
								regstr[k-1] = -1
							} else {
								regstr[k] = i - 1
							}
						}
						for y := k + 1; y < j; y++ {
							regstr[y] = -1
						}
						for y := 0; y < len; y++ {
							regstr[j+y] = i + y
						}
						j += len
						i += len
						find = true

						fmt.Printf("find %d, %d, %v\n", i, j, regstr)
						break
					}
				}
			}
		}
		return find
	}

	for k := 0; k < totalp; k++ {
		regstr[k] = -1
	}

	for {
		fmt.Printf("checking: i, j=%d, %d, %v\n", i, j, regstr)
		if j < totalp-1 && reg[j+1] == '*' {
			for {
				if i == totals {
					break
				}

				if reg[j] == '.' || reg[j] == str[i] {
					if regstr[j] == -1 {
						regstr[j] = i
					}
					regstr[j+1] = i
					i++
				} else {
					break
				}
			}
			j = j + 2
		} else if i < totals && (reg[j] == '.' || reg[j] == str[i]) {
			regstr[j] = i
			i++
			j++
		} else if !backcheck() {
			return false
		}

		if j == totalp {
			break
		}
	}

	fmt.Printf("check s=%s, p=%s, %d = %d, %d = %d, %v\n", s, p, i, totals, j, totalp, regstr)

	return i == totals && j == totalp
}
