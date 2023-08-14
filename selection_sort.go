package main

import (
	"fmt"
	"math/rand"
)

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
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
	selectionSort(arr)
	fmt.Println(arr)
}
