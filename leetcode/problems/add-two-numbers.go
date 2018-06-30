package problems

import (
	"fmt"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
	return strings.Join(ss, " -> ")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tr := result
	ltr := tr
	tl1 := l1
	tl2 := l2

	for tl1 != nil && tl2 != nil {
		if tr == nil {
			tr = &ListNode{}
			ltr.Next = tr
		}
		tr.Val = tl1.Val + tl2.Val
		tl1 = tl1.Next
		tl2 = tl2.Next
		ltr = tr
		tr = tr.Next
	}
	if tl1 != nil {
		ltr.Next = tl1
	} else if tl2 != nil {
		ltr.Next = tl2
	}
	tr = result
	upnum := 0
	for tr != nil {
		if tr.Val+upnum >= 10 {
			tr.Val = tr.Val + upnum - 10
			upnum = 1
		} else {
			tr.Val = tr.Val + upnum
			upnum = 0
		}
		if tr.Next == nil && upnum == 1 {
			tr.Next = &ListNode{
				Val: 1,
			}
			break
		}
		tr = tr.Next
	}

	return result
}
