package main

import (
	"errors"
	"fmt"
)

func main() {
	//!int
	var myInt int8 = 16
	fmt.Println(myInt)

	//uint unsigned int
	var myIntTwo uint32 = 10000
	fmt.Println(myIntTwo)

	//alias
	// INT32 = rune
	var intThree rune = 1234
	fmt.Println(intThree)

	//alias
	// INT8 = byte
	var intFour byte = 123
	fmt.Println(intFour)

	//!Float
	var mtFloat float32 = 10.66
	fmt.Println(mtFloat)

	var mtFloatTwo float64 = 10.66888282822
	fmt.Println(mtFloatTwo)

	//!String
	var myString string = "test"
	fmt.Println(myString)

	myChar := 'B' // int on ASC
	fmt.Println(myChar)

	//! Value 0 || empty
	var text string
	fmt.Println(text)

	//!bool
	firstBool := true
	fmt.Println(firstBool)

	//!Error
	var myError error
	fmt.Println(myError) // nill (default value)

	var mySecondError error = errors.New("Internal error")
	fmt.Println(mySecondError)

}
