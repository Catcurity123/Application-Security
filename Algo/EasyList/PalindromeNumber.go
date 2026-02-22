package main

import(
	"fmt"
)

func isPalindrome(x int) bool{
	if x < 0{
		return false
	}
	
	RotateNum := x
	checknum := 0
	for RotateNum != 0{
		checknum = checknum * 10 + RotateNum % 10
		RotateNum /= 10
	}
	return checknum == x
}


func main(){
	fmt.Println(isPalindrome(121))
}


/* 


*/