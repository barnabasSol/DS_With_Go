package tree

type BinaryTree struct {
	Root *Node
}

func NewBT(root int) *BinaryTree {
	return &BinaryTree{
		Root: &Node{Data: root, Count: 1},
	}
}

type Node struct {
	Count       int
	Data        int
	Left, Right *Node
}

func (bt *BinaryTree) Insert(data int) {
	bt.Root.Count++
	if bt.Root == nil {
		bt.Root = &Node{Data: data}
		return
	}
	bt.Root.insert(data)
}

func (bt *BinaryTree) Search(data int) bool {
	return bt.Root.search(data)
}

func (n *Node) search(data int) bool {
	if data == n.Data {
		return true
	}
	if data > n.Data {
		if n.Right == nil {
			return false
		}
		return n.Right.search(data)
	}
	if data < n.Data {
		if n.Left == nil {
			return false
		}
		return n.Left.search(data)
	}
	return false
}

func (n *Node) insert(data int) {
	newNode := &Node{Data: data}
	if data > n.Data {
		if n.Right == nil {
			n.Right = newNode
			return
		}
		n.Right.insert(data)
	} else if data < n.Data {
		if n.Left == nil {
			n.Left = newNode
			return
		}
		n.Left.insert(data)
	} else {
		println("duplicates arent allowed")
	}
}
