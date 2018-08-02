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

		find := false
		cj := j
		fixed := false
		for k := j - 1; k >= 0; k-- {
			fmt.Printf("bc: j=%d, cj=%d, k=%d\n", j, cj, k)
			if regstr[k] > -1 {
				if k > 0 && reg[k] == '*' {
					if reg[k-1] == '.' && reg[j] != '.' {
						for x := regstr[k]; x >= regstr[k-1]; x-- {
							fmt.Printf("ccc cj=%d, x=%d\n", cj, x)
							if str[x] == reg[j] {
								regstr[cj] = x
								if regstr[k-1] == regstr[cj] {
									regstr[k-1] = -1
									regstr[k] = -1
								} else {
									regstr[k] = regstr[cj] - 1
								}
								find = true
								break
							}
						}
						if find {
							break
						}
					} else if reg[j] == '.' || reg[k-1] == reg[j] {
						regstr[cj] = regstr[k]
						if regstr[k-1] == regstr[cj] {
							regstr[k-1] = -1
							regstr[k] = -1
						} else {
							regstr[k] = regstr[cj] - 1
						}
						fmt.Printf("right111: cj=%d, k=%d, %v\n", cj, k, regstr)
						find = true
						break
					}
					k--
				} else if reg[k] == '.' && reg[j] != '.' {
					if str[regstr[k]] == reg[j] {
						fixed = true
						regstr[cj] = regstr[k]
						regstr[k] = -1
						cj = k
					}
				} else if reg[k] == reg[j] {
					fixed = true
					regstr[cj] = regstr[k]
					regstr[k] = -1
					cj = k
					fmt.Printf("right: cj=%d, k=%d, %v\n", cj, k, regstr)
				}
			}
		}

		if find {
			i = regstr[j] + 1
			j++
			fmt.Printf("i, j = %d, %d\n", i, j)
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
