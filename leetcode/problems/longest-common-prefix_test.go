package problems

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if res != "fl" {
		t.Errorf("error: should be fl, in fact is [%s]", res)
	}
	res = longestCommonPrefix([]string{"dog", "racecar", "car"})
	if res != "" {
		t.Errorf("error: should be , in fact is [%s]", res)
	}
	res = longestCommonPrefix([]string{})
	if res != "" {
		t.Errorf("error: should be , in fact is [%s]", res)
	}
	res = longestCommonPrefix([]string{""})
	if res != "" {
		t.Errorf("error: should be , in fact is [%s]", res)
	}
}
