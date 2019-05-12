package main

import "fmt"

func main() {
	heap := MaxHeap{make([]int, 1), 0}
	for i := 0; i < 10; i++ {
		heap.insert(i)
	}
	fmt.Println(heap.data)
	fmt.Println(heap.deleteMax())
	fmt.Println(heap.data)
	heap.insert(30)
	fmt.Println(heap.data)
}

type MaxHeap struct {
	data  []int
	count int
}

func (heap *MaxHeap) getSize() int {
	return heap.count
}

func (heap *MaxHeap) isEmpty() bool {
	return heap.count == 0
}

func (heap *MaxHeap) insert(val int) {
	heap.count++
	heap.data = append(heap.data, val)
	heap.shiftUp()
}

func (heap *MaxHeap) shiftUp() {
	index := heap.count
	for index > 1 && heap.data[index/2] < heap.data[index] {
		heap.data[index], heap.data[index/2] = heap.data[index/2], heap.data[index]
		index /= 2
	}
}

func (heap *MaxHeap) deleteMax() int {
	if heap.count == 0 {
		return -1
	}
	val := heap.data[1]
	heap.data[1], heap.data[heap.count] = heap.data[heap.count], heap.data[1]
	heap.data = heap.data[:heap.count]
	heap.count--
	heap.shiftDown()
	return val
}

func (heap *MaxHeap) shiftDown() {
	index := 1
	for index*2 <= heap.count {
		childIndex := index * 2
		if childIndex+1 <= heap.count && heap.data[childIndex+1] > heap.data[childIndex] {
			childIndex++
		}
		if heap.data[index] >= heap.data[childIndex] {
			break
		}
		heap.data[index], heap.data[childIndex] = heap.data[childIndex], heap.data[index]
		index = childIndex
	}
}

func (heap *MaxHeap) getMax() int {
	if heap.count == 0 {
		return -1
	}
	return heap.data[1]
}
