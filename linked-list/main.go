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
	toPrint := l.head
	for l.linkedlist_size != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.linkedlist_size--
	}
	fmt.Println()
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

	mylist.printList()
	mylist.deleteWithValue(20)
	mylist.printList()
}
