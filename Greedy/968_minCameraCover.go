package Greedy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 每个节点的可能情况
// 它可能为空节点，叶子节点，或者非叶子节点
// 每个节点能否被摄像头照到的情况
// 它未被照到，它自己就是摄像头,它被子节点照到
// 在叶子节点安置摄像头是不合理的 所以能安置摄像头的位置一定是非叶子节点
// 采用后序遍历,每次的返回出口都是非叶子节点

func minCameraCover(root *TreeNode) int {
	number := 0
	var postOrder func(*TreeNode) int
	postOrder = func(node *TreeNode) int {
		// 定义三种状态，0表示未被监控，1表示自身就是监控，2表示被其他节点监控

		// 如果当前节点为空 则返回2 即空叶子节点可以看作是被监控（核心）
		// 如果设置为未监控 即0  则所有叶子节点都将被强迫安装摄像头
		// 如果设置未自己为摄像头 那么也会影响叶子节点的父节点的判断 空节点为摄像头 则叶子节点为状态2 那么父节点被应该安装摄像头则会向上寻找

		if node == nil {
			return 2
		}
		l := postOrder(node.Left)
		r := postOrder(node.Right)
		// l和r 都可能取得0,1,2的情况 现在分类讨论

		//情况一：l和至少有一个返回0 证明l和r至少有一个未被监控 则当前节点必须放置摄像头并返回1
		if l == 0 || r == 0 { // (l,r) --> (0,0)，(1,0) (2,0) (0,1) (0,2)
			number++
			return 1
		}
		// 情况二： 当前节点的左右孩子中至少有一个为监控 则当前节点不用放置，且当前节点属于状态3
		if l == 1 || r == 1 { // (l,r) --> (1,1) (1,2) (2,1)
			return 2
		}
		// 情况三：左右孩子节点均是被自己的子节点监控并无向上监控的能力
		// 按照我们应该向上放置摄像头的理念，当前摄像头是一个逻辑上的叶子节点 应该向上求助
		return 0
	}
	// 针对root节点也得进行处理
	rootState := postOrder(root)
	// 如果root被监控到了 无论是自己本身是摄像头还是被照顾到了 直接返回
	if rootState == 1 || rootState == 2 {
		return number
	}
	// 否则root并未被照顾到（比如一个链表形状的二叉树），还得在root放置一个
	return number + 1
}
