package main

import (
	"fmt"
	"math/rand"
)

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {

		var j = i
		for j > 0 {
			if arr[j-1] <= arr[j] {
				break
			}

			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
}

func main() {
	var arr = make([]int, 100)

	var limit = 1000
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(limit) - limit/2
	}

	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}
