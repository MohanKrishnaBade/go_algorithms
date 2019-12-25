package Healper

import (
	"math/rand"
	"time"
)

func GenerateSlice(len int) []int {

	slice := make([]int, len, len);
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		slice[i] = rand.Intn(len)

	}

	return slice;
}