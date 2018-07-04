package problems

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	res := nblengthOfLongestSubstring("abcabcbb")
	if res != 3 {
		t.Errorf("error result: should be 3, in fact is %d", res)
	}

	res = nblengthOfLongestSubstring("bbbbb")
	if res != 1 {
		t.Errorf("error result: should be 1, in fact is %d", res)
	}

	res = nblengthOfLongestSubstring("pwwkew")
	if res != 3 {
		t.Errorf("error result: should be 3, in fact is %d", res)
	}

	res = nblengthOfLongestSubstring("au")
	if res != 2 {
		t.Errorf("error result: should be 2, in fact is %d", res)
	}

	res = nblengthOfLongestSubstring("a")
	if res != 1 {
		t.Errorf("error result: should be 1, in fact is %d", res)
	}

	res = lengthOfLongestSubstring("")
	if res != 0 {
		t.Errorf("error result: should be 0, in fact is %d", res)
	}

	res = nblengthOfLongestSubstring("dvdf")
	if res != 3 {
		t.Errorf("error result: should be 3, in fact is %d", res)
	}
}
