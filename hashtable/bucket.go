package hashtable

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{Key: key}
		temp := b.head
		b.head = newNode
		b.head.next = temp
	} else {
		println("already exists nigger")
	}
}

func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}
	if b.head.Key == key {
		b.head = b.head.next
		return
	}
	prev := b.head
	current := b.head.next
	for current != nil {
		if current.Key == key {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}

}
func (b bucket) search(key string) bool {
	for b.head != nil {
		if b.head.Key == key {
			return true
		}
		b.head = b.head.next
	}
	return false
}

func (b bucket) dump() {
	for b.head != nil {
		print(b.head.Key, " ")
		b.head = b.head.next
	}
}

type bucketNode struct {
	Key  string
	next *bucketNode
}
