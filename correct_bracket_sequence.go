package main

import "fmt"

func isOpen(c rune) bool {
	return c == '[' || c == '(' || c == '{'
}

func matching(a, b rune) bool {
	return a == '[' && b == ']' || a == '(' && b == ')' || a == '{' && b == '}'
}

func isValid(s string) bool {
	var stack = make([]rune, 0)

	for _, c := range s {
		if isOpen(c) {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		var top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !matching(top, c) {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[]"))
	fmt.Println(isValid("[()]"))
	fmt.Println(isValid("[({})]"))
	fmt.Println(isValid("[({})][]"))
	fmt.Println(isValid("[]]({})][]"))
}
