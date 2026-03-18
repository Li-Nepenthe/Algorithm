package LinkedList

type MyLinkedList struct {
	dummy *LinkedList
	size  int
}

type LinkedList struct {
	val  int
	next *LinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &LinkedList{
			next: nil,
		},
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this.size <= index || index < 0 {
		return -1
	}
	p := this.dummy.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	newHead := &LinkedList{
		val: val,
	}
	p := this.dummy
	for i := 0; i < index; i++ {
		p = p.next
	}
	newHead.next = p.next
	p.next = newHead
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	p := this.dummy
	for i := 0; i < index; i++ {
		p = p.next
	}
	p.next = p.next.next
	this.size--
}
