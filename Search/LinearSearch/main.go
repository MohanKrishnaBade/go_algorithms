package main

import (
	"fmt"
	"../../../../go/go_algorithms/Helper"
)

func main() {
	info := Healper.StdInputInfo{
		Message: Healper.RADOM_ARRAY_LENGTH_MESSAGE,
	}

	Healper.StandardInput(&info)
	items := Healper.GenerateSlice(info.Int2)

	info.Message = Healper.LUCKY_NUMBER_MESSAGE
	Healper.StandardInput(&info)

	fmt.Println(linearSearch(items, info.Int2))
}

func linearSearch(data []int, key int) bool {

	for _, item := range data {
		if item == key {
			return true;
		}
	}
	return false;
}
