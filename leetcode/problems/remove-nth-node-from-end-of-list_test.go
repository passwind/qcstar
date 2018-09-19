package problems

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	nodes := makeListNode([]int{1, 2, 3, 4, 5})
	t.Logf("nodes is %s", nodes)
	res := removeNthFromEnd(nodes, 2)
	t.Logf("res is %s", res)

	nodes = makeListNode([]int{1})
	t.Logf("nodes is %s", nodes)
	res = removeNthFromEnd(nodes, 1)
	t.Logf("res is %s", res)

	nodes = makeListNode([]int{})
	t.Logf("nodes is %s", nodes)
	res = removeNthFromEnd(nodes, 0)
	t.Logf("res is %s", res)

	nodes = makeListNode([]int{1, 2, 3, 4, 5})
	t.Logf("nodes is %s", nodes)
	res = removeNthFromEnd(nodes, 5)
	t.Logf("res is %s", res)

	nodes = makeListNode([]int{1, 2, 3, 4, 5})
	t.Logf("nodes is %s", nodes)
	res = removeNthFromEnd(nodes, 1)
	t.Logf("res is %s", res)
}
