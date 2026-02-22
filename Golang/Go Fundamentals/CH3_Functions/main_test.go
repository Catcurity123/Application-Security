package main

import (
	"fmt"
	"testing"
)

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

func Test_l3(t *testing.T) {
	type testCase struct {
		tier     string
		expected int
	}

	runCases := []testCase{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
	}

	submitCases := append(runCases, []testCase{
		{"invalid", 0},
		{"", 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMonthlyPrice(test.tier)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func Test_l4(t *testing.T) {
	type testCase struct {
		costPerSend  int
		numLastMonth int
		numThisMonth int
		expected     int
	}

	runCases := []testCase{
		{2, 89, 102, 26},
		{2, 98, 104, 12},
	}

	submitCases := append(runCases, []testCase{
		{3, 50, 40, -30},
		{3, 60, 60, 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := monthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
		_ = getBillForMonth(0, 0)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}