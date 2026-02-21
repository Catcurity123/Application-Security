package main

import (
	"fmt"
)

func main(){
	l1()
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