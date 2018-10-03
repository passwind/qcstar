# [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/description/)

## 10-3 20ms 63.04%

![generate-parentheses-1.png](generate-parentheses-1.png)

## 范例 12ms

```golang
func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(s string, left, right int)
	backtrack = func(s string, left, right int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}

		if left < n {
			backtrack(s+"(", left+1, right)
		}
		if right < left {
			backtrack(s+")", left, right+1)
		}
	}

	backtrack("", 0, 0)
	return res
}
```
