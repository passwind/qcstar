package problems

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := head.Next
	c := head
	var t *ListNode

	for c != nil && c.Next != nil {
		n := c.Next
		c.Next = n.Next
		n.Next = c
		if t != nil {
			t.Next = n
		}
		t = c
		c = c.Next
	}

	return h
}
