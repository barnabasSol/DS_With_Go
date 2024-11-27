package linkedlist

import "fmt"

type Node struct {
	Data int
	next *Node
}

type linkedList struct {
	head   *Node
	length int
}

func New() *linkedList {
	return &linkedList{}
}

func (l *linkedList) Prepend(n *Node) {
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

func (l *linkedList) Append(n *Node) {
	if l.head == nil {
		l.head = n
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = n
	l.length++
}

func (dl linkedList) Find(value int) bool {
	for dl.head != nil {
		if dl.head.Data == value {
			return true
		}
		dl.head = dl.head.next
	}
	return false
}

func (l *linkedList) Delete(Data int) {
	if l.head == nil {
		return
	}
	if l.head.Data == Data {
		l.head = l.head.next
		l.length--
		return
	}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current.Data == Data {
			prev.next = current.next
			l.length--
			return
		}
		prev = current
		current = current.next
	}
}

func (l linkedList) Dump() {
	for l.head != nil {
		fmt.Print(l.head.Data, " ")
		l.head = l.head.next
	}
}
