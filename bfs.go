package main

import "fmt"

type queue struct {
	a []*node
}

func (q *queue) Push(v *node) {
	if v == nil {
		return
	}

	q.a = append(q.a, v)
}

func (q *queue) Pop() *node {
	var top = q.a[0]
	q.a = q.a[1:]

	return top
}

func (q *queue) Len() int {
	return len(q.a)
}

func bfs(root *node) []int {
	q := queue{}
	q.Push(root)

	var answer = make([]int, 0)
	for q.Len() > 0 {
		var top = q.Pop()
		q.Push(top.Left)
		q.Push(top.Right)

		answer = append(answer, top.Val)
	}

	return answer
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

	fmt.Println(bfs(root))
}
