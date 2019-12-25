package main

import (
	"fmt"
)

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(80, items))
}

func binarySearch(key int, data []int) bool {

	var (
		length = len(data);
		median = length / 2;
		value  = data[median]
	)
	if key == value {
		return true;
	} else if (key < value && length > 1) {
		return binarySearch(key, data[0:median])
	} else if (key > value && length > 1) {
		return binarySearch(key, data[median:])
	} else {
		return false;
	}
}
