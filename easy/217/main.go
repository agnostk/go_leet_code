package main

import "github.com/agnostk/go_leet_code/tester"

func main() {
	testSet := tester.TestSuite{
		Tests: map[int]tester.TestCase{
			1: {
				Args:     []int{1, 2, 3, 1},
				Expected: true,
			},
			2: {
				Args:     []int{1, 2, 3, 4},
				Expected: false,
			},
			3: {
				Args:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
				Expected: true,
			},
		},
	}
	tester.RunTestSuite(testSet, func(i interface{}) interface{} {
		return containsDuplicate(i.([]int))
	})
}

func containsDuplicate(nums []int) bool {
	var set map[int]bool = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !set[nums[i]] {
			set[nums[i]] = true
		}
	}
	return len(set) != len(nums)
}
