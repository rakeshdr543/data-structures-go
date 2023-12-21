package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (node *Node) Insert(i int) {

	if i > node.data {
		if node.right == nil {
			node.right = &Node{data: i}
		} else {
			node.right.Insert(i)
		}
	} else if i < node.data {
		if node.left == nil {
			node.left = &Node{data: i}
		} else {
			node.left.Insert(i)
		}
	}

}

func (node *Node) Search(i int) bool {
	if node == nil {
		return false
	}
	if i > node.data {
		return node.right.Search(i)
	} else if i < node.data {
		return node.left.Search(i)
	}
	return true
}

func main() {
	myTree := Node{data: 10}
	myTree.Insert(5)
	myTree.Insert(20)
	myTree.Insert(30)
	fmt.Println(myTree.Search(10))
	fmt.Println(myTree)
}
