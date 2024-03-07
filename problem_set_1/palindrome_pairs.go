package main

import (
	"fmt"
)

// isPalindrome checks if a string is palindrome
func isPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// finds all pairs of indices
func palindromePairs(words []string) [][]int {
	var result [][]int
	//iterate over all pairs of words
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			//avoid comparing words to itself
			if i != j {
				concatenated := words[i] + words[j]
				if isPalindrome(concatenated) {
					pair := []int{i, j}
					result = append(result, pair)
				}
			}
		}
	}

	return result
}

func main() {

	inputWords := []string{"bat", "tab", "cat", "rat", "tar"}
	result := palindromePairs(inputWords)
	fmt.Println(result)
}
