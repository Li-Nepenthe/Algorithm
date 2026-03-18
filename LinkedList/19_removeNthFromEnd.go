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
	// 找寻倒数n+1节点
	for i := 0; i < n && prev.Next != nil; i++ {
		prev = prev.Next
	}
	curr := dummy
	for prev.Next != nil {
		curr = curr.Next
		prev = prev.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}
