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
		t.Errorf("error: should be false, s=\"%s\", p=\"%s\"", "aabbbbbbb", "....*")
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
}
