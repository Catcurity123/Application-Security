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

/*
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
/*
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
	*/

func Test_l5(t *testing.T){
	type testCase struct {
		email    string
		username string
		domain   string
	}

	runCases := []testCase{
		{"drogon@dragonstone.com", "drogon", "dragonstone.com"},
		{"rhaenyra@targaryen.com", "rhaenyra", "targaryen.com"},
	}

	submitCases := append(runCases, []testCase{
		{"viserys@kingslanding.com", "viserys", "kingslanding.com"},
		{"aegon@stormsend.com", "aegon", "stormsend.com"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

	for _, test := range testCases {
		username, domain := splitEmail(test.email)
		if username != test.username || domain != test.domain {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.email, test.username, test.domain, username, domain)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.email, test.username, test.domain, username, domain)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func Test_l6(t *testing.T){
		type testCase struct {
		productID      string
		quantity       int
		accountBalance float64
		expected_1     bool
		expected_2     float64
	}

	runCases := []testCase{
		{"1", 2, 226.95, true, 223.95},
		{"2", 25, 459, true, 402.75},
		{"3", 7, 1185.2, false, 1185.2},
		{"4", 5, 0, false, 0},
		{"5", 50, 195, true, 70},
	}

	submitCases := append(runCases, []testCase{
		{"6", 0, 100, true, 100},
		{"7", 1, 210.24, false, 210.24},
		{"8", 55, 24.5, false, 24.5},
		{"9", 1, 999.99, true, 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output_1, output_2 := placeOrder(
			test.productID,
			test.quantity,
			test.accountBalance,
		)
		if output_1 != test.expected_1 || output_2 != test.expected_2 {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %.2f)
Expecting:  (%v, %.2f)
Actual:     (%v, %.2f)
Fail
`, test.productID, test.quantity, test.accountBalance, test.expected_1, test.expected_2, output_1, output_2)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %.2f)
Expecting:  (%v, %.2f)
Actual:     (%v, %.2f)
Pass
`, test.productID, test.quantity, test.accountBalance, test.expected_1, test.expected_2, output_1, output_2)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}