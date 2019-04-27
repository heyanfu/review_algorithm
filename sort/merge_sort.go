package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 10}
	mergeSort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func mergeSort(a []int) []int {
	length := len(a)
	if length <= 1 {
		return a
	}
	mid := length / 2
	leftArr := mergeSort(a[:mid])
	rightArr := mergeSort(a[mid:])
	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) (res []int) {
	leftIndex, rightIndex := 0, 0
	leftLength := len(leftArr)
	rightLength := len(rightArr)
	for leftIndex < leftLength && rightIndex < rightLength {
		if leftArr[leftIndex] < rightArr[rightIndex] {
			res = append(res, leftArr[leftIndex])
			leftIndex++
		} else {
			res = append(res, rightArr[rightIndex])
			rightIndex++
		}
	}
	res = append(res, leftArr[leftIndex:]...)
	res = append(res, rightArr[rightIndex:]...)

	return
}
