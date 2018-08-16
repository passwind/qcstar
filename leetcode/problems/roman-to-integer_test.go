package problems

import "testing"

func Test_romanToInt(t *testing.T) {
	res := romanToInt("III")
	if res != 3 {
		t.Errorf("error: should be 3, in fact [%d]", res)
	}

	res = romanToInt("IV")
	if res != 4 {
		t.Errorf("error: should be 4, in fact [%d]", res)
	}

	res = romanToInt("IX")
	if res != 9 {
		t.Errorf("error: should be 9, in fact [%d]", res)
	}

	res = romanToInt("LVIII")
	if res != 58 {
		t.Errorf("error: should be 58, in fact [%d]", res)
	}

	res = romanToInt("MCMXCIV")
	if res != 1994 {
		t.Errorf("error: should be 1994, in fact [%d]", res)
	}

	res = romanToInt("I")
	if res != 1 {
		t.Errorf("error: should be 1, in fact [%d]", res)
	}

	res = romanToInt("MMMCMXCIX")
	if res != 3999 {
		t.Errorf("error: should be 3999, in fact [%d]", res)
	}

	// 1. 6 - VI, error VIII
	res = romanToInt("VI")
	if res != 6 {
		t.Errorf("error: should be 6, in fact [%d]", res)
	}

	res = romanToInt("V")
	if res != 5 {
		t.Errorf("error: should be 5, in fact [%d]", res)
	}

	res = romanToInt("X")
	if res != 10 {
		t.Errorf("error: should be 10, in fact [%d]", res)
	}

	res = romanToInt("L")
	if res != 50 {
		t.Errorf("error: should be 50, in fact [%d]", res)
	}

	res = romanToInt("C")
	if res != 100 {
		t.Errorf("error: should be 100, in fact [%d]", res)
	}

	res = romanToInt("D")
	if res != 500 {
		t.Errorf("error: should be 500, in fact [%d]", res)
	}

	res = romanToInt("M")
	if res != 1000 {
		t.Errorf("error: should be 1000, in fact [%d]", res)
	}

	res = romanToInt("MDL")
	if res != 1550 {
		t.Errorf("error: should be 1550, in fact [%d]", res)
	}

	res = romanToInt("MML")
	if res != 2050 {
		t.Errorf("error: should be 2050, in fact [%d]", res)
	}

}
