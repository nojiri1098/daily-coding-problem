package problem001

import (
	"github.com/thoas/go-funk"
)

// This problem was recently asked by Google.
//
// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// numbersから数値を2つ選び、それを足しあわせたとき、kになる組み合わせがあるかどうかをたしかめる
//
// Bonus: Can you do this in one pass?
// できれば、一度の走査で判定したい

func Run(numbers []int, k int) bool {
	var diffs []int

	for _, n := range numbers {
		diff := k - n

		if funk.Contains(diffs, diff) {
			return true
		}

		diffs = append(diffs, n)
	}

	return false
}
