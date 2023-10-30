package bst

type Node struct {
	Key   int
	Right *Node
	Left  *Node
}

type BinaryTree struct {
	Root *Node
}

func NewTree(key int) *BinaryTree {
	return &BinaryTree{
		Root: &Node{
			Key: key,
		},
	}
}

func (bt *BinaryTree) Insert(root *Node, key int) {
	n := &Node{Key: key}
	if bt.Root == nil {
		bt.Root = n
	} else if root.Key < n.Key {
		if root.Right == nil {
			root.Right = n
		} else {
			bt.Insert(root.Right, key)
		}
	} else {
		if root.Left == nil {
			root.Left = n
		} else {
			bt.Insert(root.Left, key)
		}
	}
}

var count int

func (bt *BinaryTree) Search(root *Node, key int) (bool, int) {
	count++
	if root == nil {
		return false, 0
	}
	if key == root.Key {
		return true, count
	} else if root.Key < key {
		return bt.Search(root.Right, key)
	} else if root.Key > key {
		return bt.Search(root.Left, key)
	}
	return false, 0
}
