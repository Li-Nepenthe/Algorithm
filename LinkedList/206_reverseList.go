package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	// 制造一个nil指针
	var prev *ListNode
	// 指向当前头节点
	curr := head
	for curr != nil {
		// 首先保存当前指针指向的节点 无论是否为空
		next := curr.Next
		// 断开链表，指向新节点
		curr.Next = prev
		// 更新当前prev 使得下一个指向新节点
		prev = curr
		// 指针沿原链表后移
		curr = next
	}
	return prev
}
