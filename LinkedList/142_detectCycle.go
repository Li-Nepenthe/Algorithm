package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 假设从表头到入环的长度为a，从环入口到相遇点的距离为b，从相遇点到环入口的节点为c
// 则环长C = b+c
// slow = 1 fast = 2  假设slow走了n圈 fast走了m圈 相遇 m > n
// L_Slow = a+nC+b  L_fast = a+mC+b 且2*L_slow = L_fast
// a+mC+b = 2a+2nC+2b --> a = (m-2n)C-b --> a+b = (m-2n)C
// 由于a+b为正数故m-2n >= 1 移项得 a = (m-2n-1)*C + C-b
// 其中(m-2n-1)*C为圈数 C-b为相遇点到环入口的距离
// 故从相遇点出发和从表头出发 到表头走了a到环入口时 相遇点再走了(m-2n-1)*C圈后再走C-b即c后也到了环入口

func detectCycle(head *ListNode) *ListNode {

	return nil
}
