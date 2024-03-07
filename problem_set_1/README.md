Problem Set: Palindrome Pairs 
Problem Description
Given a list of unique words, the task is to find all pairs of distinct indices (i, j) in the list so that the concatenation of the two words, i.e., words[i] + words[j], forms a palindrome.

Solution Overview
The solution involves iterating through all pairs of words in the given list. For each pair, the words are concatenated, and the resulting string is checked if it's a palindrome. If it is, the pair of indices is added to the result. The process ensures that words are not compared to self and handle cases where the concatenation form palindromes.

Instructions to Run the Code
1.Ensure you have Go installed on your system.
2.Copy the provided code into a Go source file (e.g., main.go).
3.Run the program using the go run command.