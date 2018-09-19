package problems

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	var head, c, n *ListNode
	for c1 != nil || c2 != nil {
		if c1 == nil {
			if head == nil {
				return c2
			}
			c.Next = c2
			break
		} else if c2 == nil {
			if head == nil {
				return c1
			}
			c.Next = c1
			break
		} else {
			if c1.Val < c2.Val {
				n = c1
				c1 = c1.Next
			} else {
				n = c2
				c2 = c2.Next
			}
			if head == nil {
				head = n
			}
			if c == nil {
				c = head
			} else {
				c.Next = n
				c = n
			}
		}
	}
	return head
}
