// Problem Statement#
// Given a string, find the length of the longest substring, which has all distinct characters.

// Example 1:
// Input: String="aabccbb"
// Output: 3
// Explanation: The longest substring with distinct characters is "abc".

// Example 2:
// Input: String="abbbb"
// Output: 2
// Explanation: The longest substring with distinct characters is "ab".

// Example 3:
// Input: String="abccde"
// Output: 3
// Explanation: Longest substrings with distinct characters are "abc" & "cde".

package slidingwindows

import (
	"testing"

	"github.com/BrimmingDev/grokking-the-coding-interview/common"
)

// O(n) time complexity where n is the length of string, O(1) space complexity because there is at most 26 chars in the hash map
func LongestDistinctSubstring(str string) int {
	seen := make(map[string]int)
	start := 0
	maxLength := 0

	for end := 0; end < len(str); end++ {
		rightChar := string(str[end])

		seen[rightChar] += 1

		for seen[rightChar] > 1 {
			leftChar := string(str[start])
			seen[leftChar] -= 1
			start++
		}

		maxLength = common.MaxInt(maxLength, end-start+1)
	}

	return maxLength
}

func TestLongestDistinct_Example1(t *testing.T) {
	const expected = 3
	got := LongestDistinctSubstring("aabccbb")

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestLongestDistinct_Example2(t *testing.T) {
	const expected = 2
	got := LongestDistinctSubstring("abbbb")

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestLongestDistinct_Example3(t *testing.T) {
	const expected = 3
	got := LongestDistinctSubstring("abccde")

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}
