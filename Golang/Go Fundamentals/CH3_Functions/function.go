package main

import (
	"fmt"
)

func main(){
	l1()
	l2()
}

func l1(){
	// For function's parameters, it can be s1 string, s2 string, or s1, s2 string.
	// Parameters of the same type can be grouped together
	concat := func(s1, s2 string) string{
		return s1 + s2
	}

	testFunc := func(s1, s2 string) {
		fmt.Println(concat(s1, s2))
	}

	testFunc("Lane,", " happy birthday!")
	testFunc("Zuck,", " hope that Metaverse thing works out")
	testFunc("Go", " is fantastic")
}

func l2(){
	printCostReport := func(costCalculator func(string) int, message string){
		cost := costCalculator(message)
		fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
		fmt.Println()
	}

	printReports := func(intro, body, outro string){
		printCostReport(func(m string) int{
			return len(m) * 2
		}, intro)
		
		printCostReport(func(m string) int{
			return len(m) * 3
		}, body)
		
		printCostReport(func(m string) int{
			return len(m) * 4
		}, outro)
	}

	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func getMonthlyPrice(tier string) int {
	returnPennies := 0
	basicPrice := 100
	premiumPrice := 150
	enterprisePrice := 500
	
	switch tier {
	case "basic":
		returnPennies = basicPrice * 100
	case "premium":
		returnPennies = premiumPrice * 100
	case "enterprise":
		returnPennies = enterprisePrice * 100
	}
	return returnPennies
}
