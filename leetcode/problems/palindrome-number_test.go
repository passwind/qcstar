package problems

import (
	"math"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	res := isPalindrome(121)
	if !res {
		t.Errorf("error: should be true [121]")
	}

	res = isPalindrome(-121)
	if res {
		t.Errorf("error: should be false [-121]")
	}

	res = isPalindrome(10)
	if res {
		t.Errorf("error: should be false [10]")
	}

	res = isPalindrome(0)
	if !res {
		t.Errorf("error: should be true [0]")
	}

	res = isPalindrome(math.MaxInt32)
	if res {
		t.Errorf("error: should be false [%d]", math.MaxInt32)
	}

	res = isPalindrome(math.MinInt32)
	if res {
		t.Errorf("error: should be false [%d]", math.MinInt32)
	}

	res = isPalindrome(99999999999)
	if !res {
		t.Errorf("error: should be true [99999999999]")
	}
}
