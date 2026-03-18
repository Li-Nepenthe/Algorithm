package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路：由于两个链表的前驱不等长，所以强行构造等长链表即A+B
// 当A的指针走完A时 跳转到B继续走  B的指针走完时跳到A继续走
// 二者一定能在相交处相遇 这个相交处可以是真的相交点，也可以是链表末尾（不相交）

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		// 只有当 pA 走到 nil 之后才切换到 headB
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
