package problems

import "testing"

func Test_mergeTwoLists(t *testing.T) {
	// 1->2->4, 1->3->4
	// 输出：1->1->2->3->4->4
	l1 := makeListNode([]int{1, 2, 4})
	l2 := makeListNode([]int{1, 3, 4})
	res := mergeTwoLists(l1, l2)
	if res.String() != "1->1->2->3->4->4" {
		t.Errorf("error: should be [1->1->2->3->4->4], in fact: %s", res)
	}

	l1 = nil
	l2 = makeListNode([]int{1, 3, 4})
	res = mergeTwoLists(l1, l2)
	if res.String() != "1->3->4" {
		t.Errorf("error: should be [1->3->4], in fact: %s", res)
	}

	l1 = makeListNode([]int{1, 3, 4})
	l2 = nil
	res = mergeTwoLists(l1, l2)
	if res.String() != "1->3->4" {
		t.Errorf("error: should be [1->3->4], in fact: %s", res)
	}

	l1 = nil
	l2 = nil
	res = mergeTwoLists(l1, l2)
	if res.String() != "" {
		t.Errorf("error: should be [], in fact: %s", res)
	}

	l1 = makeListNode([]int{1, 3, 4})
	l2 = makeListNode([]int{100})
	res = mergeTwoLists(l1, l2)
	if res.String() != "1->3->4->100" {
		t.Errorf("error: should be [1->3->4->100], in fact: %s", res)
	}

	l1 = makeListNode([]int{2, 3, 4})
	l2 = makeListNode([]int{1})
	res = mergeTwoLists(l1, l2)
	if res.String() != "1->2->3->4" {
		t.Errorf("error: should be [1->2->3->4], in fact: %s", res)
	}
}
