package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	// fast指针先走n步 因为都是从伪头节点开始的 所以同步n
	for i := 0; i < n && prev.Next != nil; i++ {
		prev = prev.Next
	}
	curr := dummy
	// 循环结束curr指针指向倒数第n+1个节点即被删节点的前驱
	for prev.Next != nil {
		curr = curr.Next
		prev = prev.Next
	}
	// 删除倒数第n个节点
	curr.Next = curr.Next.Next
	return dummy.Next
}
