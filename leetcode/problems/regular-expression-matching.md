# [正则表达式匹配](https://leetcode-cn.com/problems/regular-expression-matching/description/)

***执行用时 4ms***

![提交记录](./regular-expression-matching.png)

## 范例 4ms

```golang
func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	var dp [][]bool
	var t []bool
	for i := 0; i <= slen; i++ {
		t = make([]bool, plen+1)
		dp = append(dp, t)
	}
	for i := 0; i <= slen; i++ {
		for j := 0; j <= plen; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			} else if i == 0 {
				dp[i][j] = ((j-1)%2 == 1 && p[j-1] == '*' && dp[i][j-2])
				continue
			} else if j == 0 {
				dp[i][j] = false
				continue
			}
			if p[j-1] != '*' {
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '.') && dp[i-1][j-1]
			} else {
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i-1][j]
				}
				if dp[i][j-2] == true {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[slen][plen]
}
```

范例算法的思路是直接用每一个s的字符去和p对比，几个关键点：

1. s的每个字符构成行，p的每个字符构成列
1. 为了比较空字符串，增加了一个行、列
1. 全匹配的状态必然是最后一个map元素为`true`
1. 如果`dp[i-1][j-1]`为`false`，则`dp[i][j]`为`false`
1. 如果是s的某个字符已经匹配上了，那么在当前行匹配以`*`结束的规则时，即便匹配结果为`false`，也应为`true`，如`aabcc`匹配`aabd*c*`，当匹配`b`和`d*`时，虽然为`false`，但仍应改为`true`。这就是如下代码的作用

```golang
if dp[i][j-2] == true {
	dp[i][j] = true
}
```
