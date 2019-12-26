package main

import (
	"../../Helper"
	"fmt"
	"runtime"
)

func main() {
	info := Healper.StdInputInfo{
		Message: Healper.RADOM_ARRAY_LENGTH_MESSAGE,
	}

	Healper.StandardInput(&info)

	slice := Healper.GenerateSlice(info.Int2)

	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	fmt.Println(runtime.NumCPU())

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
