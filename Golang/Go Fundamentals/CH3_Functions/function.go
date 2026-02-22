package main

import (
	"fmt"
)

var shouldConnectToDB = true
var shouldConnectToPaymentProvider = true

func main(){
	defineFuncAsVar()
	annonymousFunction()
	deferKeywordInVar()
}


/// L2
func defineFuncAsVar(){
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


////L3
func annonymousFunction(){
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



//// L4
func deferKeywordInVar(){
	test(true, true)
	test(false, true)
	test(true, false)
	test(false, false)
}

func bootup() {
		defer fmt.Println("TEXTIO BOOTUP DONE")
		ok := connectToDB()
		if !ok {
			return
		}

		ok = connectToPaymentProvider()
		if !ok{
			return
		}
		fmt.Println("All systems ready!")
}

func connectToDB() bool {
		fmt.Println("Connecting to database...")
		if shouldConnectToDB {
			fmt.Println("Connected!")
			return true
		}
		fmt.Println("Connection failed")
		return false
}

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

func test(dbSuccess, paymentSuccess bool){
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}


/// L5
func splitEmail(email string) (string, string){
	username, domain := "", ""
	for i, r := range email {
		if r == '@'{
			username = email[:i]
			domain = email[i + 1:]
			break
		}
	}
	return username, domain
}

/// L6
func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	orderStock := amountInStock(productID)
	if quantity <= orderStock && accountBalance >= calcPrice(productID, quantity) {
		return true, accountBalance - calcPrice(productID, quantity)
	} else {
		return false, accountBalance
	}
}

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}