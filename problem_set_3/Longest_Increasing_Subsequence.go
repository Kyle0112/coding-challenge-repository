package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] represents the length of the longest increasing subsequence ending at nums[i]
	dp := make([]int, len(nums))
	dp[0] = 1 // Base case: at least the element itself forms an increasing subsequence

	// Initialize the dp array with 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		// For each element, check all the elements before it to find the longest increasing subsequence
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	// Find the maximum value in the dp array
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func main() {
	// Example usage
	inputNumbers := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(inputNumbers)
	fmt.Println(result) // Output: 4
}