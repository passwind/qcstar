package problems

import "testing"

func equalStrArray(v, t []string) bool {
	existedElement := func(s string, tt []string) bool {
		for _, e := range tt {
			if s == e {
				return true
			}
		}
		return false
	}
	for _, s := range v {
		if !existedElement(s, t) {
			return false
		}
	}
	for _, s := range t {
		if !existedElement(s, v) {
			return false
		}
	}
	return true
}
func Test_letterCombinations(t *testing.T) {
	// 	输入："23"
	// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
	res := letterCombinations("23")
	if !equalStrArray(res, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		t.Errorf("error: should be %v, in fact is - %v", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, res)
	}
	// t.Logf("%v", res)
}
