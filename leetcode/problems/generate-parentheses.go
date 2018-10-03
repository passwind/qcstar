package problems

func generateParenthesis(n int) []string {
	res := []string{}
	var gp func(s string, l, r int)
	gp = func(s string, l, r int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}
		if l < n {
			gp(s+"(", l+1, r)
		}
		if r < l {
			gp(s+")", l, r+1)
		}
	}
	gp("", 0, 0)
	return res
}

func generateParenthesis1(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}

	pmap := make(map[string]int)
	addC := func(ts string) {
		if _, ok := pmap[ts]; !ok {
			pmap[ts] = 1
		}
	}

	for i := 1; i <= n/2; i++ {
		if i == 1 {
			tres := generateParenthesis(n - i)
			for _, s := range tres {
				addC("()" + s)
				addC(s + "()")
				addC("(" + s + ")")
			}
		} else if i == n-i {
			tres := generateParenthesis(i)
			for _, s := range tres {
				for _, t := range tres {
					addC(s + t)
				}
			}
		} else {
			tres1 := generateParenthesis(i)
			tres2 := generateParenthesis(n - i)
			for _, s := range tres1 {
				for _, t := range tres2 {
					addC(s + t)
					addC(t + s)
				}
			}
		}
	}

	res := []string{}
	for k := range pmap {
		res = append(res, k)
	}

	return res
}
