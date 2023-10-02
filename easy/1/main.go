package main

import (
	"github.com/agnostk/go_leet_code/tester"
)

type ArgsType struct {
	nums   []int
	target int
}

func main() {
	testSuite := tester.TestSuite{
		Tests: map[int]tester.TestCase{
			1: {
				Args: ArgsType{
					nums:   []int{2, 7, 11, 15},
					target: 9,
				},
				Expected: []int{0, 1},
			},
			2: {
				Args: ArgsType{
					nums:   []int{3, 2, 4},
					target: 6,
				},
				Expected: []int{1, 2},
			},
			3: {
				Args: ArgsType{
					nums:   []int{3, 3},
					target: 6,
				},
				Expected: []int{0, 1},
			},
		},
	}
	tester.RunTestSuite(testSuite, func(i interface{}) interface{} {
		return twoSum(i.(ArgsType).nums, i.(ArgsType).target)
	})
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0}
}
