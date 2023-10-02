package main

import "github.com/agnostk/go_leet_code/tester"

type ArgsType struct {
	nums []int
}

func main() {
	testSet := tester.TestSuite{
		Tests: map[int]tester.TestCase{
			1: {
				Args: ArgsType{
					nums: []int{1, 2, 3, 1},
				},
				Expected: true,
			},
			2: {
				Args: ArgsType{
					nums: []int{1, 2, 3, 4},
				},
				Expected: false,
			},
			3: {
				Args: ArgsType{
					nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
				},
				Expected: true,
			},
		},
	}
	tester.RunTestSuite(testSet, func(i interface{}) interface{} {
		return containsDuplicate(i.(ArgsType).nums)
	})
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !set[nums[i]] {
			set[nums[i]] = true
		}
	}
	return len(set) != len(nums)
}
