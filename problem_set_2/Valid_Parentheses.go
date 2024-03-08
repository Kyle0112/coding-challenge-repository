package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (char == ')' && top != '(') || (char == ']' && top != '[') || (char == '}' && top != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	// Example usage
	inputString := "()[]{}"
	result := isValid(inputString)
	fmt.Println(result) // Output: true

	// Additional test cases
	inputString = "([)]"
	result = isValid(inputString)
	fmt.Println(result) // Output: false

	inputString = "{[()]}"
	result = isValid(inputString)
	fmt.Println(result) // Output: true
}
