package problems

import "testing"

func Test_myAtoi(t *testing.T) {
	res := myAtoi("42")
	if res != 42 {
		t.Errorf("error: should be 42, in fact is %d", res)
	}

	res = myAtoi("   -42")
	if res != -42 {
		t.Errorf("error: should be -42, in fact is %d", res)
	}

	res = myAtoi("4193 with words")
	if res != 4193 {
		t.Errorf("error: should be 4193, in fact is %d", res)
	}

	res = myAtoi("words and 987")
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}

	res = myAtoi("-91283472332")
	if res != -2147483648 {
		t.Errorf("error: should be -2147483648, in fact is %d", res)
	}

	res = myAtoi("")
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}

	res = myAtoi("                     ")
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}

	// 1
	res = myAtoi("+-2")
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}

	// 2
	res = myAtoi("-2147483648")
	if res != -2147483648 {
		t.Errorf("error: should be -2147483648, in fact is %d", res)
	}

	res = myAtoi("2147483647")
	if res != 2147483647 {
		t.Errorf("error: should be 2147483647, in fact is %d", res)
	}
}
