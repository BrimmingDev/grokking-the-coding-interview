package slidingwindows

import (
	"math"
	"testing"

	"github.com/BrimmingDev/grokking-the-coding-interview/common"
)

// Problem Statement#
// Given an array of positive numbers and a positive number ‘k,’ find the maximum sum of any contiguous subarray of size ‘k’.

// Example 1:

// Input: [2, 1, 5, 1, 3, 2], k=3
// Output: 9
// Explanation: Subarray with maximum sum is [5, 1, 3].
// Example 2:

// Input: [2, 3, 4, 1, 5], k=2
// Output: 7
// Explanation: Subarray with maximum sum is [3, 4].

func MaxSumSubarraySizeK(k int, arr []int) int {
	windowStart := 0
	maxSum := math.MinInt
	currSum := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		currSum += arr[windowEnd]

		if windowEnd >= k-1 {
			maxSum = common.MaxInt(maxSum, currSum)
			currSum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

func TestExample1(t *testing.T) {
	const expected = 9
	got := MaxSumSubarraySizeK(3, []int{2, 1, 5, 1, 3, 2})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestExample2(t *testing.T) {
	const expected = 7
	got := MaxSumSubarraySizeK(2, []int{2, 3, 4, 1, 5})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}
