package main

import "fmt"

type node struct {
	Val   int
	Left  *node
	Right *node
}

func leafWithTargetSumma(node *node, currentSum, targetSum int) bool {
	if node.Left == nil && node.Right == nil {
		return currentSum+node.Val == targetSum
	}

	var leftHave bool
	if node.Left != nil {
		leftHave = leafWithTargetSumma(node.Left, currentSum+node.Val, targetSum)
	}

	var rightHave bool
	if node.Right != nil {
		rightHave = leafWithTargetSumma(node.Right, currentSum+node.Val, targetSum)
	}

	return leftHave || rightHave
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

	if root == nil {
		fmt.Println(false)
	}

	fmt.Println(leafWithTargetSumma(root, 0, 12))
	fmt.Println(leafWithTargetSumma(root, 0, 18))
	fmt.Println(leafWithTargetSumma(root, 0, 5))
}
