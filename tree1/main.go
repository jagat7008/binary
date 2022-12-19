package main

import "fmt"

var count int

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) insert(k int) {
	count++
	if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		return n.right.Search(k)
	} else if n.key > k {
		return n.left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{key: 100}
	tree.insert(50)
	tree.insert(204)
	tree.insert(102)
	tree.insert(23)
	tree.insert(26)
	tree.insert(56)
	fmt.Println(tree.Search(56))
	fmt.Println(count)

}
