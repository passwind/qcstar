package problems

import "testing"

func Test_intToRoman(t *testing.T) {
	res := intToRoman(3)
	if res != "III" {
		t.Errorf("error: should be III, in fact [%s]", res)
	}

	res = intToRoman(4)
	if res != "IV" {
		t.Errorf("error: should be IV, in fact [%s]", res)
	}

	res = intToRoman(9)
	if res != "IX" {
		t.Errorf("error: should be IX, in fact [%s]", res)
	}

	res = intToRoman(58)
	if res != "LVIII" {
		t.Errorf("error: should be LVIII, in fact [%s]", res)
	}

	res = intToRoman(1994)
	if res != "MCMXCIV" {
		t.Errorf("error: should be MCMXCIV, in fact [%s]", res)
	}
}
