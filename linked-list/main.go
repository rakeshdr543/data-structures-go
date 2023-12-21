package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head            *node
	linkedlist_size int
}

func (l *linkedlist) prepend(node *node) {
	node.next = l.head
	l.head = node
	l.linkedlist_size++
}

func (l linkedlist) printList() {
	cur := l.head
	for cur != nil {
		fmt.Printf("%d ", cur.data)
		cur = cur.next
	}
	fmt.Println()
}

func (l *linkedlist) reverse() {
	if l.linkedlist_size == 0 {
		return
	}

	var prev *node
	cur := l.head
	for cur != nil {

		temp := cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}
	l.head = prev
}

func (l linkedlist) middle() node {
	slow, fast := l.head, l.head.next
	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return *slow
}

func (l *linkedlist) deleteWithValue(value int) {
	if l.linkedlist_size == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.linkedlist_size--
		return
	}

	prevNode := l.head
	for prevNode.next.data != value {
		if prevNode.next.next == nil {
			return
		}
		prevNode = prevNode.next
	}

	prevNode.next = prevNode.next.next
	l.linkedlist_size--
}

func main() {
	mylist := linkedlist{}
	mylist.prepend(&node{data: 10})
	mylist.prepend(&node{data: 20})
	mylist.prepend(&node{data: 30})
	mylist.prepend(&node{data: 40})
	mylist.prepend(&node{data: 50})
	mylist.prepend(&node{data: 60})

	mylist.printList()
	mylist.reverse()
	mylist.printList()
	fmt.Println(mylist.middle())
}
