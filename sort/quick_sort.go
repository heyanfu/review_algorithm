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

func quickSort2(a []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	i, j, val := leftIndex+1, rightIndex, a[leftIndex]
	for {
		for i <= rightIndex && a[i] < val {
			i++
		}
		for j >= leftIndex+1 && a[j] > val {
			j--
		}
		if i > j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	a[leftIndex], a[j] = a[j], a[leftIndex]
	quickSort2(a, leftIndex, j-1)
	quickSort2(a, j+1, rightIndex)
}

func quickSort3(a []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	l, r, i := leftIndex, rightIndex+1, leftIndex+1
	val := a[leftIndex]
	for i < r {
		if a[i] < val {
			a[i], a[l+1] = a[l+1], a[i]
			l++
			i++
		} else if a[i] > val {
			a[i], a[r-1] = a[r-1], a[i]
			r--
		} else {
			i++
		}
	}
	a[leftIndex], a[l] = a[l], a[leftIndex]
	quickSort3(a, leftIndex, l-1)
	quickSort3(a, r, rightIndex)
}
