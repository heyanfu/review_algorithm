package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 10}
	heapSort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func heapSort(a []int) {
	length := len(a)
	for i := (length - 1) / 2; i >= 0; i-- {
		shiftDown(a, length, i)
	}

	for i := length - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		shiftDown(a, i, 0)
	}
}

func shiftDown(a []int, length, index int) {
	for index*2+1 < length {
		childIndex := 2*index + 1
		if childIndex+1 < length && a[childIndex+1] > a[childIndex] {
			childIndex++
		}
		if a[index] >= a[childIndex] {
			break
		}
		a[index], a[childIndex] = a[childIndex], a[index]
		index = childIndex
	}
}
