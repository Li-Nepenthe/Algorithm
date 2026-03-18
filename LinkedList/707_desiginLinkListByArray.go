package LinkedList

type MyLinkedList2 struct {
	value []int //存储数据
	nexts []int //存储下一个节点的数组下标
	head  int   //虚拟头节点下标
	free  []int
	size  int
}

func Constructor2() MyLinkedList2 {
	my := MyLinkedList2{
		value: make([]int, 1, 1001),
		nexts: make([]int, 1, 1001),
		head:  0,
	}
	my.nexts[0] = -1
	return my
}

func (this *MyLinkedList2) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	// 找到index对应的数组下标
	pre := this.head
	for i := 0; i <= index; i++ {
		pre = this.nexts[pre]
	}
	return this.value[pre]
}

func (this *MyLinkedList2) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	prev := this.head
	for i := 0; i < index; i++ {
		prev = this.nexts[prev] //寻找前驱节点
	}
	newNodeIndex := len(this.value)
	this.value = append(this.value, val)              // 向value数组末尾追加元素
	this.nexts = append(this.nexts, this.nexts[prev]) // newNode.next = prev.next
	this.nexts[prev] = newNodeIndex                   // pre.next = newNode
	this.size++
}

func (this *MyLinkedList2) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList2) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList2) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	// 1. 定位前驱
	preIdx := this.head
	for i := 0; i < index; i++ {
		preIdx = this.nexts[preIdx]
	}

	// 2. 找到要被踢出的那个下标
	targetIdx := this.nexts[preIdx]

	// 3. 逻辑跳过
	this.nexts[preIdx] = this.nexts[targetIdx]

	// 4. 【高级操作】回收这个下标到 free 链表（可选）
	// 将来 AddAtIndex 时可以先看 free 里有没有值，有就复用
	// this.nexts[targetIdx] = this.free
	// this.free = targetIdx

	// 5. 必须更新元数据
	this.size--
}
