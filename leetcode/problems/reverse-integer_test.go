package problems

import "testing"

func Test_reverse(t *testing.T) {
	res := reverse(321)
	if res != 123 {
		t.Errorf("error: should be 123, in fact is %d", res)
	}
	res = reverse(-123)
	if res != -321 {
		t.Errorf("error: should be -321, in fact is %d", res)
	}
	res = reverse(120)
	if res != 21 {
		t.Errorf("error: should be 21, in fact is %d", res)
	}
	res = reverse(2147483650)
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}

	res = reverse(-2147483650)
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}
	res = reverse(0)
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}
	res = reverse(1534236469)
	if res != 0 {
		t.Errorf("error: should be 0, in fact is %d", res)
	}
}
