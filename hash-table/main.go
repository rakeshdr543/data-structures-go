package main

import "fmt"

func generateHash(s string) int {
	count := 0
	for _, ch := range s {
		count += int(ch)
	}

	return count % 5
}

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

type HashTable struct {
	items [5]*LinkedList
}

func (h *HashTable) Insert(s string) {
	index := generateHash(s)
	ll := h.items[index]
	if ll == nil {
		ll = &LinkedList{}
		h.items[index] = ll
	}
	ll.Insert(s)
}

func (h *HashTable) Search(s string) bool {
	index := generateHash(s)
	ll := h.items[index]
	if ll == nil {
		ll = &LinkedList{}
		h.items[index] = ll

	}
	return ll.Search(s)
}

func (h *HashTable) Remove(s string) {
	index := generateHash(s)
	ll := h.items[index]
	if ll == nil {
		ll = &LinkedList{}
		h.items[index] = ll

	}
	ll.Remove(s)
}

func (l *LinkedList) Insert(s string) {
	newNode := &Node{
		data: s,
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) Remove(s string) {
	head := l.head
	if head == nil {
		return
	}
	if head.data == s {
		l.head = nil
		return
	}
	var prev *Node
	cur := head
	for cur.data != s {
		prev = cur
		cur = cur.next
		if cur == nil {
			return
		}
	}
	prev.next = cur.next
}

func (l LinkedList) Search(s string) bool {
	head := l.head
	if head == nil {
		return false
	}

	for head != nil {
		if head.data == s {
			return true
		}
		head = head.next
	}
	return false
}

func main() {
	hashTable := HashTable{}
	hashTable.Insert("rahul")
	hashTable.Insert("raki")
	hashTable.Insert("rami")
	fmt.Println(hashTable.Search("Sandy"))
	fmt.Println(hashTable.Search("raki"))
	hashTable.Remove("raki")
	fmt.Println(hashTable.Search("raki"))
}
