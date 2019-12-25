package main

import (
	"fmt"
	"../../../go_algorithms/Helper"
)

func main() {

	info := Healper.StdInputInfo{
		Message: Healper.RADOM_ARRAY_LENGTH_MESSAGE,
	}

	Healper.StandardInput(&info)
	items := Healper.GenerateSlice(info.Int2)

	info.Message = Healper.LUCKY_NUMBER_MESSAGE
	Healper.StandardInput(&info)

	fmt.Println(binarySearch(info.Int2, items))

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
