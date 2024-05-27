package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array3)

	//Slices

	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1)
	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	// []int
	fmt.Println(reflect.TypeOf(array1))
	// [5]int

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array1[1] = 99
	fmt.Println(slice2)

	/// Internals arrays

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length //10
	fmt.Println(cap(slice3)) // capacity //11

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length //11
	fmt.Println(cap(slice3)) // capacity //11

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length //12
	fmt.Println(cap(slice3)) // capacity //24 (doubled the size internal, because are full)

	slice4 := make([]float32, 5)

	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length //5
	fmt.Println(cap(slice4)) // capacity //5 (If I don't set the capacity, it's the same as length)
}
