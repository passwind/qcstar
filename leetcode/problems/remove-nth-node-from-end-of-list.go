package problems

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	r := head
	nmap := make(map[int]*ListNode)
	i := 0
	for r != nil {
		nmap[i] = r
		i++
		r = r.Next
	}
	c := nmap[i-n]
	if c == head {
		return head.Next
	}

	p := nmap[i-n-1]
	p.Next = c.Next

	return head
}
