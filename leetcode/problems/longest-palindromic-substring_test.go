package problems

import "testing"

func Test_longestPalindrome(t *testing.T) {
	res := longestPalindrome("babad")
	if res != "bab" {
		t.Errorf("error: should be [bab], in fact is [%s]", res)
	}
	res = longestPalindrome("aba")
	if res != "aba" {
		t.Errorf("error: should be [aba], in fact is [%s]", res)
	}
	res = longestPalindrome("cbbd")
	if res != "bb" {
		t.Errorf("error: should be [bb], in fact is [%s]", res)
	}
	res = longestPalindrome("b")
	if res != "b" {
		t.Errorf("error: should be [b], in fact is [%s]", res)
	}
	res = longestPalindrome("bb")
	if res != "bb" {
		t.Errorf("error: should be [bb], in fact is [%s]", res)
	}
	res = longestPalindrome("bbb")
	if res != "bbb" {
		t.Errorf("error: should be [bbb], in fact is [%s]", res)
	}
	res = longestPalindrome("bbbb")
	if res != "bbbb" {
		t.Errorf("error: should be [bbbb], in fact is [%s]", res)
	}
	res = longestPalindrome("")
	if res != "" {
		t.Errorf("error: should be [], in fact is [%s]", res)
	}
	res = longestPalindrome("abcba")
	if res != "abcba" {
		t.Errorf("error: should be [abcba], in fact is [%s]", res)
	}
	res = longestPalindrome("abbbbabcbabbbbaxxx")
	if res != "abbbbabcbabbbba" {
		t.Errorf("error: should be [abbbbabcbabbbba], in fact is [%s]", res)
	}
	res = longestPalindrome("yyyyyyyabbbbabcbabbbbaxxx")
	if res != "abbbbabcbabbbba" {
		t.Errorf("error: should be [abbbbabcbabbbba], in fact is [%s]", res)
	}
	res = longestPalindrome("bananas")
	if res != "anana" {
		t.Errorf("error: should be [anana], in fact is [%s]", res)
	}
}
