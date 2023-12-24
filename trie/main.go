package main

import "fmt"

type Node struct {
	isLast   bool
	children [26]*Node
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(s string) {
	cur := t.root

	for _, ch := range s {
		index := int(ch) - int('a')
		if cur.children[index] == nil {
			cur.children[index] = &Node{}
		}
		cur = cur.children[index]
	}
	cur.isLast = true
}

func (t *Trie) Search(s string) bool {
	cur := t.root

	for _, ch := range s {
		index := int(ch) - int('a')
		if cur.children[index] == nil {
			return false
		}
		cur = cur.children[index]
	}
	return cur != nil && cur.isLast
}

func main() {

	trie := Trie{
		root: &Node{},
	}
	trie.Insert("rakesh")
	fmt.Println(trie.Search("rakesh"))
	fmt.Println(trie.Search("ramyy"))

}
