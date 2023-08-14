package main

import "fmt"

func solution(arr []int) int {
	var maxSub, current int
	for i := 0; i < len(arr); i++ {
		current += arr[i]
		current = max(current, 0)
		maxSub = max(maxSub, current)
	}

	return maxSub
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(solution(arr))
}
