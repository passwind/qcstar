package problems

import "testing"

func Test_swapPairs(t *testing.T) {
	// 给定 1->2->3->4, 你应该返回 2->1->4->3
	l1 := makeListNode([]int{1, 2, 3, 4})
	l2 := makeListNode([]int{2, 1, 4, 3})
	res := swapPairs(l1)
	if res.String() != l2.String() {
		t.Errorf("error: should be [%s], in fact: %s", l2, res)
	}

	l1 = makeListNode([]int{})
	l2 = makeListNode([]int{})
	res = swapPairs(l1)
	if res.String() != l2.String() {
		t.Errorf("error: should be [%s], in fact: %s", l2, res)
	}

	l1 = makeListNode([]int{1, 2, 3})
	l2 = makeListNode([]int{2, 1, 3})
	res = swapPairs(l1)
	if res.String() != l2.String() {
		t.Errorf("error: should be [%s], in fact: %s", l2, res)
	}

	l1 = makeListNode([]int{1, 2})
	l2 = makeListNode([]int{2, 1})
	res = swapPairs(l1)
	if res.String() != l2.String() {
		t.Errorf("error: should be [%s], in fact: %s", l2, res)
	}
}
