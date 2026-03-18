package LinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

//func removeElements(head *ListNode, val int) *ListNode {
//	//先处理头节点一直满足val的情况
//	for head != nil && head.Val == val {
//		head = head.Next
//	}
//	if head == nil {
//		return nil
//	}
//	var p = head
//	for p.Next != nil {
//		if p.Next.Val == val {
//			p.Next = p.Next.Next
//		} else {
//			p = p.Next
//		}
//
//	}
//	return head
//}

// 哨兵版本
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
