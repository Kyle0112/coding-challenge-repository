Problem Set 2: Valid Parentheses

Problem Description
Given a string containing just the characters '(', ')', '{', '}', '[', and ']', determine if the input string is valid. An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

Examples
Input: ()[]{}
Output: True
Input: ([)]
Output: False

Solution Overview
To solve this problem, i need to ensure that the open brackets are closed by the same type of brackets and are closed in the correct order. I achieved this by using a stack data structure to keep track of the opening brackets encountered so far. I iterate through the characters of the input string and perform the necessary checks to determine the validity of the parentheses.

Instructions to Run the Code
1. Ensure you have Go installed on your system.
2. Copy the provided code into a Go source file (e.g., main.go).
3. Run the program using the go run command.

Testing Your Solution
- Test your solution with various strings to verify its correctness.
- Ensure that open brackets are closed by the same type of brackets and in the correct order.