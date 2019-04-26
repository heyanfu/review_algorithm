package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 10}
	insertionSort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func insertionSort(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}
