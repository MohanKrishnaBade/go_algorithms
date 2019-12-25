package main

import (
	"fmt"
)

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearSearch(items, 90))
}

func linearSearch(data []int, key int) bool {

	for _, item := range data {
		if item == key {
			return true;
		}
	}
	return false;
}
