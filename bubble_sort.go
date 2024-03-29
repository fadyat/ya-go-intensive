package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func main() {
	var arr = make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(1000)
	}

	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}
