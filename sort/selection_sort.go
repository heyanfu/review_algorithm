package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 10}
	selectionSort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func selectionSort(a []int) {
	length := len(a)
	for i := 0; i < length - 1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}
		a[minIndex], a[i] = a[i], a[minIndex]
	}
}
