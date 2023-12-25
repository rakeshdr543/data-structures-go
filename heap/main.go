package main

import "fmt"

type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(data int) {

	h.items = append(h.items, data)
	h.HeapifyInsertUp(len(h.items) - 1)

}

func (h *MaxHeap) HeapifyInsertUp(index int) {
	for h.items[index] > h.items[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}

}

func (h *MaxHeap) HeapifyExtract() {
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.maxHeapifyDown(0)
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	length := len(h.items) - 1
	if length == 0 {
		return
	}

	toCompare := 0
	for left(index) <= length {
		l, r := left(index), right(index)
		toCompare = l

		if r <= length && h.items[l] < h.items[r] {
			toCompare = r
		}
		if h.items[toCompare] > h.items[index] {
			h.swap(index, toCompare)
			index = toCompare
		} else {
			return
		}
	}

}

func (h *MaxHeap) swap(index1, index2 int) {

	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

func parent(index int) int {
	return index / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func main() {
	maxPQ := &MaxHeap{}
	items := []int{10, 40, 20, 30, 50, 100}
	for _, v := range items {
		maxPQ.Insert(v)
		fmt.Println(maxPQ)
	}

	for i := 0; i < len(items); i++ {
		maxPQ.HeapifyExtract()
		fmt.Println(maxPQ)
	}

}
