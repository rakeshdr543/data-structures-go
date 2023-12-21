package main

import "fmt"

type stack struct {
	items []int
}

type queue struct {
	items []int
}

func (q *queue) enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *queue) dequeue() {
	q.items = q.items[1:len(q.items)]
}

func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *stack) pop() int {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func main() {
	// myStack := stack{
	// 	items: []int{1, 2},
	// }

	// myStack.push(5)
	// myStack.push(7)
	// myStack.push(9)
	// myStack.pop()

	// fmt.Println(myStack.items)

	myQueue := queue{
		items: []int{2, 4},
	}

	myQueue.enqueue(6)
	myQueue.enqueue(8)
	myQueue.dequeue()

	fmt.Println(myQueue.items)
}
