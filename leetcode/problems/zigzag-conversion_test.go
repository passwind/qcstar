package problems

import "testing"

func Test_convert(t *testing.T) {
	res := convert("PAYPALISHIRING", 3)
	if res != "PAHNAPLSIIGYIR" {
		t.Errorf("error: should be [PAHNAPLSIIGYIR], infact [%s]", res)
	}

	res = convert("PAYPALISHIRING", 4)
	if res != "PINALSIGYAHRPI" {
		t.Errorf("error: should be [PINALSIGYAHRPI], infact [%s]", res)
	}

	res = convert("P", 4)
	if res != "P" {
		t.Errorf("error: should be [P], infact [%s]", res)
	}

	res = convert("", 4)
	if res != "" {
		t.Errorf("error: should be [], infact [%s]", res)
	}

	res = convert("AB", 1)
	if res != "AB" {
		t.Errorf("error: should be [AB], infact [%s]", res)
	}

	res = convert("ABC", 1)
	if res != "ABC" {
		t.Errorf("error: should be [ABC], infact [%s]", res)
	}

	res = convert("ABCD", 2)
	if res != "ACBD" {
		t.Errorf("error: should be [ACBD], infact [%s]", res)
	}
}
