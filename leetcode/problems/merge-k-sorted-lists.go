package problems

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	var m2l = func(l1 *ListNode, l2 *ListNode) *ListNode {
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

	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = m2l(lists[i], lists[i+k])
		}
		n = k
	}

	return lists[0]
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var h, c *ListNode
	var ll []*ListNode
	for _, l := range lists {
		if l != nil {
			ll = append(ll, l)
		}
	}
	for len(ll) > 0 {
		mv, mi := ll[0].Val, 0
		for i, cc := range ll[1:] {
			if cc.Val < mv {
				mv = cc.Val
				mi = i + 1
			}
		}
		if c == nil {
			c = ll[mi]
			if h == nil {
				h = c
			}
		} else {
			c.Next = ll[mi]
			c = c.Next
		}
		var nl []*ListNode
		for i, cc := range ll {
			if i == mi {
				if cc.Next != nil {
					nl = append(nl, cc.Next)
				}
			} else {
				nl = append(nl, cc)
			}
		}
		c.Next = nil
		if len(nl) == 0 {
			break
		}
		ll = []*ListNode{}
		ll = append(ll, nl...)
	}
	return h
}
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var m2l = func(l1 *ListNode, l2 *ListNode) *ListNode {
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

	ll := lists[0]
	for _, l := range lists[1:] {
		ll = m2l(ll, l)
	}
	return ll
}
