package problems

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	v1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	v2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res := addTwoNumbers(v1, v2)
	t.Logf("res is %s", res)

	v1 = &ListNode{}
	v2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res = addTwoNumbers(v1, v2)
	t.Logf("res is %s", res)

	v1 = &ListNode{
		Val: 5,
	}
	v2 = &ListNode{
		Val: 5,
	}
	res = addTwoNumbers(v1, v2)
	t.Logf("res is %s", res)

	v1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
	v2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res = addTwoNumbers(v1, v2)
	t.Logf("res is %s", res)
}
