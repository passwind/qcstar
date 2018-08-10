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

	res = intToRoman(1)
	if res != "I" {
		t.Errorf("error: should be I, in fact [%s]", res)
	}

	res = intToRoman(3999)
	if res != "MMMCMXCIX" {
		t.Errorf("error: should be MMMCMXCIX, in fact [%s]", res)
	}

	// 1. 6 - VI, error VIII
	res = intToRoman(6)
	if res != "VI" {
		t.Errorf("error: should be VI, in fact [%s]", res)
	}

	res = intToRoman(5)
	if res != "V" {
		t.Errorf("error: should be V, in fact [%s]", res)
	}

	res = intToRoman(10)
	if res != "X" {
		t.Errorf("error: should be X, in fact [%s]", res)
	}

	res = intToRoman(50)
	if res != "L" {
		t.Errorf("error: should be L, in fact [%s]", res)
	}

	res = intToRoman(100)
	if res != "C" {
		t.Errorf("error: should be C, in fact [%s]", res)
	}

	res = intToRoman(500)
	if res != "D" {
		t.Errorf("error: should be D, in fact [%s]", res)
	}

	res = intToRoman(1000)
	if res != "M" {
		t.Errorf("error: should be M, in fact [%s]", res)
	}

	res = intToRoman(1550)
	if res != "MDL" {
		t.Errorf("error: should be MDL, in fact [%s]", res)
	}

	res = intToRoman(2050)
	if res != "MML" {
		t.Errorf("error: should be MML, in fact [%s]", res)
	}

}
