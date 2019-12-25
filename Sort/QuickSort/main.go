package main

import (
	"../Healper"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(quickSort(Healper.GenerateSlice(4)))

}

//it still needs some work to be done.
func quickSortNewVersion(a []int) []int {
	if len(a) <2 {
		return a
	}

	left, right := make([]int, 0), make([]int, 0);
	pivotValue := a[0]

	for _, value := range a {
		if value <= pivotValue {
			left = append(left, value);
		} else if value > pivotValue {
			right = append(right, value)

		}
	}

	return append(append(quickSort(left), []int{pivotValue}...), quickSort(right)...)
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
