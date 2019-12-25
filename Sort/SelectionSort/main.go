package main

import (
	"../Healper"
	"fmt"
)

func main() {
	array := Healper.GenerateSlice(10)
	selectionSort(array)
	fmt.Println(array)

}

func selectionSort(items []int) {

	n := len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
