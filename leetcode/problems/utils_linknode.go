package problems

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	ll := l
	ss := []string{}
	for ll != nil {
		ss = append(ss, fmt.Sprintf("%d", ll.Val))
		ll = ll.Next
	}
	return strings.Join(ss, "->")
}

func makeListNode(vals []int) *ListNode {
	var h, c *ListNode
	for _, v := range vals {
		if h == nil {
			h = &ListNode{
				Val: v,
			}
			c = h
		} else {
			n := &ListNode{
				Val: v,
			}
			c.Next = n
			c = n
		}
	}
	return h
}
