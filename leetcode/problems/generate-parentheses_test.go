package problems

import "testing"

func Test_generateParenthesis(t *testing.T) {
	res := generateParenthesis(3)
	t.Logf("res is [3] %v", res)
	res = generateParenthesis(0)
	t.Logf("res is [0] %v", res)
	res = generateParenthesis(1)
	t.Logf("res is [1] %v", res)
	res = generateParenthesis(2)
	t.Logf("res is [2] %v", res)
	res = generateParenthesis(4)
	t.Logf("res is [4] - [%d] %v", len(res), res)
	res = generateParenthesis(5)
	t.Logf("res is [4] - [%d] %v", len(res), res)
}
