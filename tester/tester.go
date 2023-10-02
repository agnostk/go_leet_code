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
			if len(expected.([]int)) != len(actual.([]int)) {
				success = false
			} else {
				for i := 0; i < len(expected.([]int)); i++ {
					if expected.([]int)[i] != actual.([]int)[i] {
						success = false
						break
					}
				}
			}
		} else if reflect.TypeOf(expected).String() == "[][]string" {
			defer func() {
				if r := recover(); r != nil {
					success = false
				}
			}()
			if len(expected.([][]string)) != len(actual.([][]string)) {
				success = false
			} else {
			outerLoop:
				for i := 0; i < len(expected.([][]string)); i++ {
					if len(expected.([][]string)[i]) != len(actual.([][]string)[i]) {
						success = false
						break outerLoop
					}
					for j := 0; j < len(expected.([][]string)[i]); j++ {
						if expected.([][]string)[i][j] != actual.([][]string)[i][j] {
							success = false
							break outerLoop
						}
					}
				}
			}
		} else {
			success = expected == actual
		}
		if success {
			fmt.Printf("Test %v | OK\n", testIndex)
		} else {
			fmt.Printf("Test %v | (E) %v | %v (A)\n", testIndex, expected, actual)
		}
	}
}
