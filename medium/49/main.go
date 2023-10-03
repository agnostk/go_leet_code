package main

import "github.com/agnostk/go_leet_code/tester"

type ArgsType struct {
	strs []string
}

func main() {
	testSuite := tester.TestSuite{
		Tests: map[int]tester.TestCase{
			1: {
				Args: ArgsType{
					strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
				},
				Expected: [][]string{
					{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"},
				},
			},
			2: {
				Args: ArgsType{
					strs: []string{""},
				},
				Expected: [][]string{
					{""},
				},
			},
			3: {
				Args: ArgsType{
					strs: []string{"a"},
				},
				Expected: [][]string{
					{"a"},
				},
			},
		},
	}
	tester.RunTestSuite(testSuite, func(i interface{}) interface{} {
		return groupAnagrams(i.(ArgsType).strs)
	})
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]uint][]string)
	for _, s := range strs {
		count := [26]uint{}
		for _, c := range s {
			count[c-'a'] += 1
		}
		hashMap[count] = append(hashMap[count], s)
	}
	groupedAnagrams := [][]string{}
	for _, group := range hashMap {
		groupedAnagrams = append(groupedAnagrams, group)
	}
	return groupedAnagrams
}
