package main

import "fmt"

type User struct {
	name    string
	age     uint8
	address Address
}

type Address struct {
	locale string
	number uint8
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.age = 18
	user1.name = "Leo"
	fmt.Println(user1)

	user2 := User{
		"leo",
		18,
		Address{
			"internet",
			10,
		},
	}
	fmt.Println(user2)

	user3 := User{
		name: "leo",
	}

	fmt.Println(user3)

}
