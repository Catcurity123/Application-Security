package main

import (
	"fmt"
)

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "vip-pro":
		fallthrough
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func l1() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// We can write this conventionally
	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	// OR use initial statement of an If block
	message := "This is the message"
	if messageLength := len(message); messageLength <= maxMessageLen{
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not set")
	}
}

func l2(){
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}

func l3(){
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

	finalCost = bulkMessageCost
	if isPremiumUser{
		finalCost = finalCost - finalCost * discountRate
	}

	if accountBalance >= finalCost{
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}


func main(){
	l1()
	l2()
	l3()
}


