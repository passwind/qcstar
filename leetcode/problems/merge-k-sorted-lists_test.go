package problems

import "testing"

func Test_mergeKLists(t *testing.T) {
	in := []*ListNode{
		makeListNode([]int{1, 4, 5}),
		makeListNode([]int{1, 3, 4}),
		makeListNode([]int{2, 6}),
	}
	res := mergeKLists(in)
	if res.String() != "1->1->2->3->4->4->5->6" {
		t.Errorf("error: should be [1->1->2->3->4->4->5->6], in fact is [%s]", res)
	}
	in = []*ListNode{
		makeListNode([]int{}),
	}
	res = mergeKLists(in)
	if res.String() != "" {
		t.Errorf("error: should be [], in fact is [%s]", res)
	}
	in = []*ListNode{
		makeListNode([]int{}),
		makeListNode([]int{1, 3, 4}),
		makeListNode([]int{}),
	}
	res = mergeKLists(in)
	if res.String() != "1->3->4" {
		t.Errorf("error: should be [1->3->4], in fact is [%s]", res)
	}
}
