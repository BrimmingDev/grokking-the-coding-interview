// Problem Statement#
// You are visiting a farm to collect fruits. The farm has a single row of fruit trees. You will be given two baskets, and your goal is to pick as many fruits as possible to be placed in the given baskets.

// You will be given an array of characters where each character represents a fruit tree. The farm has following restrictions:

// Each basket can have only one type of fruit. There is no limit to how many fruit a basket can hold.
// You can start with any tree, but you can’t skip a tree once you have started.
// You will pick exactly one fruit from every tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.
// Write a function to return the maximum number of fruits in both baskets.

// Example 1:

// Input: Fruit=["A", "B", "C", "A", "C"]
// Output: 3
// Explanation: We can put 2 "C" in one basket and one "A" in the other from the subarray ["C", "A", "C"]
// Example 2:

// Input: Fruit=["A", "B", "C", "B", "B", "C"]
// Output: 5
// Explanation: We can put 3 "B" in one basket and two "C" in the other basket.
// This can be done if we start with the second letter: ["B", "C", "B", "B", "C"]
package slidingwindows

import (
	"testing"

	"github.com/BrimmingDev/grokking-the-coding-interview/common"
)

func FruitsIntoBaskets(fruits []string) int {
	const maxKindOfFruit = 2
	fruitsPicked := make(map[string]int)
	maxPicked := 0

	start := 0
	for end := 0; end < len(fruits); end++ {
		rightFruit := fruits[end]

		fruitsPicked[rightFruit] += 1

		for len(fruitsPicked) > maxKindOfFruit {
			leftFruit := fruits[start]
			start++

			fruitsPicked[leftFruit] -= 1

			if fruitsPicked[leftFruit] == 0 {
				delete(fruitsPicked, leftFruit)
			}
		}

		maxPicked = common.MaxInt(maxPicked, end-start+1)
	}

	return maxPicked
}

func TestFruitsIntoBaskets_Example1(t *testing.T) {
	const expected = 3
	got := FruitsIntoBaskets([]string{"A", "B", "C", "A", "C"})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}

func TestFruitsIntoBaskets_Example2(t *testing.T) {
	const expected = 5
	got := FruitsIntoBaskets([]string{"A", "B", "C", "B", "B", "C"})

	if got != expected {
		t.Errorf("Expected %d but received %d", expected, got)
	}
}
