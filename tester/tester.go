package tester

import (
	"fmt"
	"reflect"
)

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
		var success bool = true
		if reflect.TypeOf(expected).String() == "[]int" {
			for i := 0; i < len(expected.([]int)); i++ {
				if expected.([]int)[i] != actual.([]int)[i] {
					success = false
					break
				}
			}
		} else {
			success = expected != actual
		}
		if !success {
			fmt.Printf("Test %v | (E) %v | %v (A)\n", testIndex, expected, actual)
		} else {
			fmt.Printf("Test %v | OK\n", testIndex)
		}
	}
}
