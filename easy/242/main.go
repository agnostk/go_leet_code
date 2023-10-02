package main

import "github.com/agnostk/go_leet_code/tester"

type ArgsType struct {
	s string
	t string
}

func main() {
	testSuite := tester.TestSuite{
		Tests: map[int]tester.TestCase{
			1: {
				Args: ArgsType{
					s: "anagram",
					t: "nagaram",
				},
				Expected: true,
			},
			2: {
				Args: ArgsType{
					s: "rat",
					t: "car",
				},
				Expected: false,
			},
			3: {
				Args: ArgsType{
					s: "a",
					t: "ab",
				},
				Expected: false,
			},
		},
	}
	tester.RunTestSuite(testSuite, func(i interface{}) interface{} {
		return isAnagram(i.(ArgsType).s, i.(ArgsType).t)
	})
}

func isAnagram(s string, t string) bool {
	keyMap := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		keyMap[s[i]] += 1
		keyMap[t[i]] -= 1
	}
	for _, c := range keyMap {
		if c != 0 {
			return false
		}
	}
	return true
}
