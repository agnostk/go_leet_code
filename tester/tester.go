package tester

import "fmt"

type TestSuite struct {
	Tests map[int]TestCase
}

type TestCase struct {
	Args, Expected interface{}
}

type TestFunction func(interface{}) interface{}

func RunTestSuite(testSuite TestSuite, function TestFunction) {
	for testIndex, testCase := range testSuite.Tests {
		expected := testCase.Expected
		actual := function(testCase.Args)
		if expected != actual {
			fmt.Printf("Test %v | (E) %v | %v (A)\n", testIndex, expected, actual)
		} else {
			fmt.Printf("Test %v | OK\n", testIndex)
		}
	}
}
