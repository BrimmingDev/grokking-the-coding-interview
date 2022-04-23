package slidingwindows

import (
	"testing"

	"github.com/BrimmingDev/grokking-the-coding-interview/common"
)

// Problem Statement#
// Given an array of positive numbers and a positive number ‘S,’ find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’. Return 0 if no such subarray exists.

// Example 1:

// Input: [2, 1, 5, 2, 3, 2], S=7
// Output: 2
// Explanation: The smallest subarray with a sum greater than or equal to ‘7’ is [5, 2].

// Example 2:

// Input: [2, 1, 5, 2, 8], S=7
// Output: 1
// Explanation: The smallest subarray with a sum greater than or equal to ‘7’ is [8].

// Example 3:

// Input: [3, 4, 1, 1, 6], S=8
// Output: 3
// Explanation: Smallest subarrays with a sum greater than or equal to ‘8’ are [3, 4, 1] or [1, 1, 6].

// O(n) time complexity, O(1) space complexity
func SmallestSubarrayWithGreaterSum(num int, arr []int) int {
	minArraySize := len(arr) + 1
	windowStart := 0
	currSum := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		currSum += arr[windowEnd]

		for currSum >= num {
			minArraySize = common.MinInt(minArraySize, windowEnd-windowStart+1)
			currSum -= arr[windowStart]
			windowStart++
		}
	}

	if minArraySize == len(arr)+1 {
		return 0
	}

	return minArraySize
}

func TestSmallestSubarrayWithGreatSum_Example1(t *testing.T) {
	const expected = 2
	got := SmallestSubarrayWithGreaterSum(7, []int{2, 1, 5, 2, 3, 2})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestSmallestSubarrayWithGreatSum_Example2(t *testing.T) {
	const expected = 1
	got := SmallestSubarrayWithGreaterSum(7, []int{2, 1, 5, 2, 8})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestSmallestSubarrayWithGreatSum_Example3(t *testing.T) {
	const expected = 3
	got := SmallestSubarrayWithGreaterSum(8, []int{3, 4, 1, 1, 6})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestSmallestSubarrayWithGreatSum_Example4(t *testing.T) {
	const expected = 0
	got := SmallestSubarrayWithGreaterSum(20, []int{3, 4, 1, 1, 6})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}
