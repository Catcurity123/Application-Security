package main

import (
	"fmt"
	"strings"
)

///L1 Struct format
type messageToSend1 struct{
	phoneNumber int
	message string
}

func structInGo(){
	//%#v provide verbose format including type and field name
	//%T provides type of variable
	type messageToSend struct{
		phoneNumber int
		message string
	}

	testMessage := messageToSend{
		phoneNumber: 12345,
		message: "This is test message",
	}

	fmt.Printf("The type of testMessage is %#v \n", testMessage)
	fmt.Printf("The type of testMessage is %T\n", testMessage)

}


///L2 Nested Struct
type messageToSend2 struct{
	message string
	sender user
	recipient user
}

type user struct{
	name string
	number int
}

func (u user) checkSend() bool {
	return strings.TrimSpace(u.name) != "" && u.number != 0
}

func canSendMessage(mToSend messageToSend2) bool {
	return mToSend.sender.checkSend() && mToSend.recipient.checkSend() 
}


///L3 Embedded Struct

/* 
Embedded structs provide a kind of data-only inheritance that can be useful at times
type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
*/

/* Embedded vs Nested
Embedded struct's field can be accessible at the top level like normal field, while nested struct we need to access each level

For example

type car struct{
	brand string
	model string
}

type truck struct{
	car
	bedSize int
}

==> We can access truck.brand and truck.model just like truck.bedSize

But if for nested struct

type messageToSend2 struct{
	message string
	sender user
	recipient user
}

type user struct{
	name string
	number int
}

We can only access the name of the user by messageToSend2.user.name but not messageToSend2.name
*/

type sender struct{
	user
	rateLimit int
}

func getSenderLog(s sender) string {
	return fmt.Sprintf(`
====================================
Sender name: %v
Sender number: %v
Sender rateLimit: %v
====================================
`, s.name, s.number, s.rateLimit)
}


/// L4 Struct Method
/*

Methods are function that is defined specifically for a struct, it is a function with a reciever.
The receiver is a special parameter that syntactically goes before the name of the function
*/

type authenticationInfo struct{
	username string
	password string
}

func (authdata authenticationInfo) getBasicAuth() string{
	return fmt.Sprintf("Authorization: Basic %s:%s", authdata.username, authdata.password)
}


/// L5 Memory Layout





func main(){
	structInGo()
}
