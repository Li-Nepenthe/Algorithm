package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func swapPairs(head *ListNode) *ListNode {
//	fakeNode := &ListNode{Next: head}
//	if head == nil {
//		return nil
//	}
//	p, q, origin := head, head.Next, fakeNode
//	for q != nil {
//		next := q.Next
//		p.Next = next
//		q.Next = p
//		origin.Next = q
//		if p.Next == nil {
//			break
//		}
//		origin = p
//		p = origin.Next
//		q = p.Next
//	}
//	return fakeNode.Next
//}

// 逻辑不变思想进行优化
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	// 确保每次都有两个节点可以交换
	for prev.Next != nil && prev.Next.Next != nil {
		p, q := prev.Next, prev.Next.Next // 分别指向要交换的节点
		p.Next = q.Next
		q.Next = p
		prev.Next = q
		prev = p
	}
	return dummy.Next
}
