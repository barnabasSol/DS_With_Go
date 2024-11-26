package dlinkedlist

import "fmt"

type Node struct {
	Data       int
	next, prev *Node
}

type dlinkedList struct {
	head   *Node
	tail   *Node
	Length int
}

func New() *dlinkedList {
	return &dlinkedList{}
}

func (dl dlinkedList) ReverseDump() {
	for dl.tail != nil {
		fmt.Print(dl.tail.Data, " ")
		dl.tail = dl.tail.prev
	}
}

func (dl dlinkedList) ForwardDump() {
	for dl.head != nil {
		fmt.Print(dl.head.Data, " ")
		dl.head = dl.head.next
	}
}
func (dl *dlinkedList) Prepend(node *Node) {
	dl.Length++
	if dl.head == nil {
		dl.head = node
		dl.tail = node
		return
	}
	temp := dl.head
	dl.head = node
	dl.head.next = temp
	temp.prev = dl.head
}

func (dl *dlinkedList) Append(node *Node) {
	dl.Length++
	if dl.head == nil {
		dl.head = node
		dl.tail = node
		return
	}
	temp := dl.head
	for temp != nil {
		if temp.next == nil {
			temp.next = node
			node.prev = temp
			dl.tail = node
			return
		}
		temp = temp.next
	}
}

func (dl *dlinkedList) Delete(value int) {
	if dl.head == nil {
		println("its empty")
		return
	}
	if dl.head.Data == value {
		dl.Length--
		if dl.head.next == nil {
			dl.tail = nil
			dl.head = nil
			return
		}
		dl.head = dl.head.next
		dl.head.prev = nil
		dl.tail = dl.head
		return
	}
	temp := dl.head.next
	for temp != nil {
		if temp.Data == value {
			if temp.next == nil {
				dl.tail = temp.prev
				dl.tail.next = nil
				temp = nil
				return
			}
			temp.next.prev = temp.prev
			temp.prev.next = temp.next
			if temp.next == nil {
				dl.tail = temp
			}
			dl.Length--
			return
		}
		temp = temp.next
	}
	println("not found")
}
