package Healper

import (
	"math/rand"
	"time"
	"fmt"
	"log"
)


func GenerateSlice(len int) []int {

	slice := make([]int, len, len);
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		slice[i] = rand.Intn(len)

	}

	return slice;
}

func StandardInput(info *StdInputInfo) {
	fmt.Print(info.Message)
	_, err := fmt.Scanf("%d", &info.Int2)

	if err != nil {
		log.Fatal(err)
	}
}
