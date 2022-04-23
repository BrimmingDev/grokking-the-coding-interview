// Problem Statement#
// Given a string, find the length of the longest substring in it with no more than K distinct characters.

// Example 1:
// Input: String="araaci", K=2
// Output: 4
// Explanation: The longest substring with no more than '2' distinct characters is "araa".

// Example 2:
// Input: String="araaci", K=1
// Output: 2
// Explanation: The longest substring with no more than '1' distinct characters is "aa".

// Example 3:
// Input: String="cbbebi", K=3
// Output: 5
// Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".

// Example 4:
// Input: String="cbbebi", K=10
// Output: 6
// Explanation: The longest substring with no more than '10' distinct characters is "cbbebi".

package slidingwindows

import (
	"testing"

	"github.com/BrimmingDev/grokking-the-coding-interview/common"
)

// O(n) time complexity where n is the length of the string, O(k) space complexity where k is the size of map we would have to store
func LongestSubstringKDistinct(str string, k int) int {
	if k >= len(str) {
		return len(str)
	}

	seen := make(map[string]int)
	maxLength := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := string(str[windowEnd])

		seen[rightChar] += 1

		for len(seen) > k {
			leftChar := string(str[windowStart])
			windowStart++

			seen[leftChar] -= 1
			if seen[leftChar] == 0 {
				delete(seen, leftChar)
			}
		}

		maxLength = common.MaxInt(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func TestLongestSubstringKDistinct_Example1(t *testing.T) {
	const expected = 4
	got := LongestSubstringKDistinct("araaci", 2)

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestLongestSubstringKDistinct_Example2(t *testing.T) {
	const expected = 2
	got := LongestSubstringKDistinct("araaci", 1)

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestLongestSubstringKDistinct_Example3(t *testing.T) {
	const expected = 5
	got := LongestSubstringKDistinct("cbbebi", 3)

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestLongestSubstringKDistinct_Example4(t *testing.T) {
	const expected = 6
	got := LongestSubstringKDistinct("cbbebi", 10)

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}
