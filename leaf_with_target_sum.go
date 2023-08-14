package main

import "fmt"

func leafWithTargetSumma(node *node, currentSum, targetSum int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val
	if node.Left == nil && node.Right == nil {
		return currentSum == targetSum
	}

	return leafWithTargetSumma(node.Left, currentSum, targetSum) || leafWithTargetSumma(node.Right, currentSum, targetSum)
}

func main() {
	root := &node{
		Val: 1,
		Left: &node{
			Val: 2,
			Left: &node{
				Val: 4,
				Left: &node{
					Val: 5,
				},
			},
		},
		Right: &node{
			Val: 3,
			Left: &node{
				Val: 4,
				Left: &node{
					Val: 10,
				},
			},
		},
	}

	fmt.Println(leafWithTargetSumma(root, 0, 12))
	fmt.Println(leafWithTargetSumma(root, 0, 18))
	fmt.Println(leafWithTargetSumma(root, 0, 5))
}
