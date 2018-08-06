package problems

import "testing"

func Test_isMatch(t *testing.T) {
	res := isMatch("aa", "a")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aa", "a")
	}

	res = isMatch("aa", "a*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aa", "a*")
	}

	res = isMatch("ab", ".*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "ab", ".*")
	}

	res = isMatch("aab", "c*a*b")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aab", "c*a*b")
	}

	res = isMatch("mississippi", "mis*is*p*.")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "mississippi", "mis*is*p*.")
	}

	res = isMatch("aab", "c.*b")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aab", "c.*b")
	}

	res = isMatch("aabbbbbbb", "c.*b")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "c.*b")
	}

	res = isMatch("aabbbbbbb", "c.*b*")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "c.*b*")
	}

	res = isMatch("aabbbbbbb", "ca*b*")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "ca*b*")
	}

	res = isMatch("", "")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "", "")
	}

	res = isMatch("", "ca*b*")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "", "ca*b*")
	}

	res = isMatch("aabbbbbbb", "")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "")
	}

	res = isMatch("aabbbbbbb", "....*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aabbbbbbb", "....*")
	}

	res = isMatch("aabbbbbbb", "...c.*")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "...c.*")
	}

	res = isMatch("aabbbbbbb", ".*.c.*")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", ".*.c.*")
	}

	res = isMatch("aabbbbbbb", "**.c.*")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "**.c.*")
	}

	// 1.
	res = isMatch("mississippi", "mis*is*ip*.")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "mississippi", "mis*is*ip*.")
	}

	res = isMatch("mississippi", "...........")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "mississippi", "...........")
	}

	res = isMatch("mississippi", "mississippi")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "mississippi", "mississippi")
	}

	// 2.
	res = isMatch("ab", ".*c")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "ab", ".*c")
	}

	// 3.
	res = isMatch("aaa", "aaaa")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aaa", "aaaa")
	}

	// 4.
	res = isMatch("aaa", "a*a")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aaa", "a*a")
	}

	res = isMatch("aaab", "a*ab*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aaab", "a*ab*")
	}

	// 5.
	res = isMatch("abbbcd", "ab*bbbcd")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aaab", "a*ab*")
	}

	// 6.
	res = isMatch("aaa", "ab*a*c*a")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "aaa", "ab*a*c*a")
	}

	// 7.
	res = isMatch("bbbba", ".*a*a")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbba", ".*a*a")
	}

	res = isMatch("bbbba", ".*a*ba")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbba", ".*a*ba")
	}

	res = isMatch("bbbba", ".*a*bba")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbba", ".*a*bba")
	}

	res = isMatch("bbbba", ".*a*bb*a")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbba", ".*a*bb*a")
	}

	res = isMatch("bbbccba", ".*a*a")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbccba", ".*a*a")
	}

	res = isMatch("bbbccba", ".*a*ccba")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbccba", ".*a*ccb*a")
	}

	res = isMatch("bbbccba", ".*a*cba")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbccba", ".*a*cba")
	}

	res = isMatch("bcbaccbc", ".*a*cba.*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bcbaccbc", ".*a*cba.*")
	}

	res = isMatch("bcbaccbc", ".*a*cbac*bc")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bcbaccbc", ".*a*cbac*bc")
	}

	// 8.
	res = isMatch("ab", ".*..")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "ab", ".*..")
	}

	// 9.
	res = isMatch("", ".*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "", ".*")
	}

	res = isMatch("", "a*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "", "a*")
	}

	res = isMatch("", "a*.*b*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "", "a*.*b*")
	}

	// 10.
	res = isMatch("", ".")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "", ".")
	}

	res = isMatch("", "a")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "", "a")
	}

	//11.
	res = isMatch("", ".*b")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "", ".*b")
	}

	// 12.
	res = isMatch("bbbaccbbbaababbac", ".b*b*.*...*.*c*.")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "bbbaccbbbaababbac", ".b*b*.*...*.*c*.")
	}

	// 13.
	res = isMatch("bccbbabcaccacbcacaa", ".*b.*c*.*.*.c*a*.c")
	if res {
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "bccbbabcaccacbcacaa", ".*b.*c*.*.*.c*a*.c")
	}

	// 14.
	res = isMatch("cabbbbcbcacbabc", ".*b.*.ab*.*b*a*c")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "cabbbbcbcacbabc", ".*b.*.ab*.*b*a*c")
	}

	// 15.
	res = isMatch("abbaaaabaabbcba", "a*.*ba.*c*..a*.a*.")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "abbaaaabaabbcba", "a*.*ba.*c*..a*.a*.")
	}
}

func Test_isMatchOne(t *testing.T) {
	res := isMatch("", ".*")
	if !res {
		t.Errorf("error: should be true, s=\"%s\", p=\"%s\"", "", ".*")
	}
}
