package main

import (
	"fmt"
	"../../Sort/Healper"
)

func main() {

	slice := Healper.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}


func bubblesort(slice []int) {
	var (
		n      = len(slice)
		sorted = false
	)

	for !sorted {

		swap := false;
		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swap = true;
			}
		}

		if !swap {
			sorted = true;
		}
	}
}
