package main

import "fmt"

func main() {
	heap := MaxIndexHeap{make([]int, 1), make([]int, 1), 0}
	for i := 0; i < 10; i++ {
		heap.insert(i, i)
	}
	fmt.Println(heap.data)
	fmt.Println(heap.indexes)
	fmt.Println("=============")
	fmt.Println(heap.deleteMax())
	fmt.Println("=============")
	fmt.Println(heap.data)
	fmt.Println(heap.indexes)
	fmt.Println("=============")
	heap.insert(9, 30)
	fmt.Println(heap.data)
	fmt.Println(heap.indexes)
}

type MaxIndexHeap struct {
	data    []int
	indexes []int
	count   int
}

func (heap *MaxIndexHeap) getSize() int {
	return heap.count
}

func (heap *MaxIndexHeap) isEmpty() bool {
	return heap.count == 0
}

func (heap *MaxIndexHeap) insert(index int, val int) {
	heap.count++

	heap.data = append(heap.data, val)
	heap.indexes = append(heap.indexes, index+1)
	heap.shiftUp()
}

func (heap *MaxIndexHeap) shiftUp() {
	index := heap.count
	for index > 1 && heap.data[heap.indexes[index/2]] < heap.data[heap.indexes[index]] {
		heap.indexes[index/2], heap.indexes[index] = heap.indexes[index], heap.indexes[index/2]
		index /= 2
	}
}

func (heap *MaxIndexHeap) deleteMax() int {
	if heap.count == 0 {
		return -1
	}
	val := heap.data[heap.indexes[1]]
	heap.indexes[1], heap.indexes[heap.count] = heap.indexes[heap.count], heap.indexes[1]
	heap.data = heap.data[:heap.count]
	heap.indexes = heap.indexes[:heap.count]
	heap.count--
	heap.shiftDown()

	return val
}

func (heap *MaxIndexHeap) shiftDown() {
	index := 1
	for index*2 <= heap.count {
		childIndex := index * 2
		if childIndex+1 <= heap.count && heap.data[heap.indexes[childIndex+1]] > heap.data[heap.indexes[childIndex]] {
			childIndex++
		}
		if heap.data[heap.indexes[index]] >= heap.data[heap.indexes[childIndex]] {
			break
		}
		heap.indexes[index], heap.indexes[childIndex] = heap.indexes[childIndex], heap.indexes[index]
		index = childIndex
	}
}

func (heap *MaxIndexHeap) getMax() int {
	if heap.count == 0 {
		return -1
	}
	return heap.data[1]
}
