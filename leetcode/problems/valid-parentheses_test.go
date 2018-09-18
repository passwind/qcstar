package problems

import "testing"

func Test_isValid(t *testing.T) {
	res := isValid("()")
	if !res {
		t.Errorf("error: should be true, in fact [%t]", res)
	}

	res = isValid("()[]{}")
	if !res {
		t.Errorf("error: should be true, in fact [%t]", res)
	}

	res = isValid("(]")
	if res {
		t.Errorf("error: should be false, in fact [%t]", res)
	}

	res = isValid("([)]")
	if res {
		t.Errorf("error: should be false, in fact [%t]", res)
	}

	res = isValid("{[]}")
	if !res {
		t.Errorf("error: should be true, in fact [%t]", res)
	}

	res = isValid("[(]")
	if res {
		t.Errorf("error: should be false, in fact [%t]", res)
	}

	res = isValid("[({}])")
	if res {
		t.Errorf("error: should be false, in fact [%t]", res)
	}

	res = isValid("[({})")
	if res {
		t.Errorf("error: should be false, in fact [%t]", res)
	}

	// 1. ]
	res = isValid("]")
	if res {
		t.Errorf("error: should be false, in fact [%t]", res)
	}
}
