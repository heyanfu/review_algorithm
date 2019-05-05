package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 10}
	quickSort(a, 0, len(a)-1)
	for _, v := range a {
		fmt.Println(v)
	}
}

func quickSort(a []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	midIndex := leftIndex
	for i := leftIndex + 1; i <= rightIndex; i++ {
		if a[i] <= a[leftIndex] {
			midIndex++
			a[i], a[midIndex] = a[midIndex], a[i]
		}
	}
	a[leftIndex], a[midIndex] = a[midIndex], a[leftIndex]
	quickSort(a, leftIndex, midIndex-1)
	quickSort(a, midIndex+1, rightIndex)
}
