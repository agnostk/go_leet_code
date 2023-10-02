package main

import "github.com/agnostk/go_leet_code/tester"

func main() {
	testSuite := tester.TestSuite{
		Tests: map[int]tester.TestCase{
			1: {
				Args:     []string{"anagram", "nagaram"},
				Expected: true,
			},
			2: {
				Args:     []string{"rat", "car"},
				Expected: false,
			},
			3: {
				Args:     []string{"a", "ab"},
				Expected: false,
			},
		},
	}
	tester.RunTestSuite(testSuite, func(i interface{}) interface{} {
		return isAnagram(i.([]string)[0], i.([]string)[1])
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
