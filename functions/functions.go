package main

import "fmt"

func main() {
	soma := add(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("text function f")
	fmt.Println(result)

	resultAdd, resultSub := calculationsMath(10, 15)
	fmt.Println(resultAdd, resultSub)
}

func add(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculationsMath(n1, n2 int8) (int8, int8) {
	add := n1 + n2
	subtraction := n1 - n2

	return add, subtraction
}
