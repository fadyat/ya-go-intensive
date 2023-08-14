package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func dfs(root *Node) int {
	if root == nil {
		return 0
	}

	return dfs(root.Left) + dfs(root.Right) + root.Val
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
				Left: &Node{
					Val: 5,
				},
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 4,
				Left: &Node{
					Val: 4,
				},
			},
		},
	}

	fmt.Println(dfs(root))
}
