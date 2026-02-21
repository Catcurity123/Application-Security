package main

import (
	"fmt"
	"unicode/utf8"
)

// warusOPeratorCantBeUsedOutsideFunction := true
var warusOPeratorCantBeUsedOutsideFunction bool //Warus Operator can not be used outside of function


func l2(){
	smsSendingLimit := 24
	costPerSMS := 2.4
	hasPermission := true
	username := "Josh"
	// We can also use normal initialization patter | var smsSendingLimit int 

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func l3(){
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)
}

func l4(){
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := float64(numMessagesFromDoris) * float64(costPerMessage)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}

func l5(){
	// Signed Integers (both posi and nega) int int8 int16 int32 int64
	// Unsigned Integers (only posi) uint uint8 uint16 uint32 uint64
	// Signed decimal float32 float64
	// Complex numbers complex64 complex128

	// The size (8, 16, 32, 64, 128) represents how many bits in memory will be used to store variable, the default int or uint is either 32 or 64 depends on the environment
	
	// Converting between types can be done like
	// temperatureFloat := 17.9
	// temperatureInt := int64(temperatureFloat)
	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

func l6(){
	fmt.Println("\nGo enforces static typing meaning variable types are known before the code runs. That means the editor and the compiler can display type error before the code is run") 
}

func l7(){
	fmt.Println("\nFor compiled, you can run a compiled program without the source code, but interpreted program need the source code so that the interpreter can be translated at runtime")
}

func l8(){
	fmt.Println("\nWe can decalre multiple variables on the sameline")
	averageOpenRate, displayMessage := .24, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}

func l9(){
	fmt.Println("\nGo programs are fairly lightweight, each program includes a small amount of extra code included in the executable binary called \"Go Runtime\". One of the purepose is to clean up unused memory at runtime")
}

func l10(){
	fmt.Println("\nConstants are declared as \"const\" it can be of any primitive types, but it cant be more complex types")
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
	fmt.Println("\nConstant can be computet as long as the computation can happen at compile time.\nconst currentTime =time.Now() is not allowed")
}

func l11(){
	// fmt.Printf(): to print a formatted string to standard out
	// fmt.Sprintf(): returns the formatted string
	// %v: prints any value in a default format, it can be used as a catchall
	//s := fmt.Sprintf("I am %v years old", 10)
	//or
	//s2 := fmt.Sprintf("I am %v years old", "way too many")
	// %s: for string
	// %d: for integer
	// %f: for float but we can add decimal points using %.2f or %.3f
	// %t: for boolean
	const name = "Saul Goodman"
	const openRate = 30.5
    msg := fmt.Sprintf("Hi %s, your open rate is %.1f%% \n", name, openRate)
	fmt.Print(msg)
}

func l12(){
	//In many programming language a "character" is a single byte. Using ASCII encoding, we can represent 128 characters with 7 bits
	//In Go, strings are just sequences of bytes; they can hold arbitrary data. Go also has a special type "rune"
	//rune is an alias for int32. This means a rune is a 32-bit integer, which is large enough to hold any Unicode code point

	//When you need to work with individual characters in a string, you should use the rune type. It breaks string up into their individual characters, and we can hold a wide variety of Unicode characters such as emojis or chinese characters
	const name = "🐻"
	//var bear rune

	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}

func l13(){
	var startup string = "Textio SMS service booting up..."
	var message string = "Sending test message"
	var confirmation string = "Message sent!"

	// don't touch below this line

	fmt.Println(startup)
	fmt.Println(message)
	fmt.Println(confirmation)
}

func l14(){
	senderName := "Syl"
	recipient := "Kaladin"
	message := "The Words, Kaladin. You have to speak the Words!"

	fmt.Printf("%s to %s: %s\n", senderName, recipient, message)
}

func l15(){
	penniesPerText := float64(2)

	//printf(%T, var) to get the type of variable
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}

func l16(){
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.10f, Is Subscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)

	// Don't touch below this line

	fmt.Println(userLog)
}

func main(){
	l2()
	l3()
	l4()
	l5()
	l6()
	l7()
	l8()
	l9()
	l10()
	l11()
	l12()
	l16()
}