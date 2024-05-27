package main

import "fmt"

func main() {
	//arithmetics
	add := 1 + 1
	sub := 1 - 1
	divide := 1 / 1
	restOfDivision := 10 % 2

	fmt.Println(add, sub, divide, restOfDivision)

	/* var number1 int16 = 10
		var number2 int32 = 200

		 Can't add different types
	addPlus= number1 + number2 */

	//Attribution

	var var1 string = "string"
	var2 := "string"

	fmt.Printf(var1, var2)

	//Relations
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Logics

	fmt.Println(true && true || true || !true)

	//unarms
	number := 10
	number++
	fmt.Println(number)
	number += 15
	fmt.Println(number)
	number--
	number -= 5
	fmt.Println(number)

	number *= 3
	number /= 2
	number %= 2
	fmt.Println(number)

	//ternary
	var text string
	if number > 5 {
		text = "more than 5"
	} else {
		text = "lesser than 5"
	}
	fmt.Println(text)
}
