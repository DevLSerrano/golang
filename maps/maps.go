package main

import "fmt"

func main() {
	fmt.Println("Maps")

	users := map[string]string{
		"name":     "Leonardo",
		"lastName": "Serrano",
	}
	fmt.Println(users)
	fmt.Println(users["name"]) // Leonardo

	users2 := map[string]map[string]string{
		"name": {
			"first":  "Leonardo",
			"second": "Serrano",
		},
	}
	fmt.Println(users2)
	fmt.Println(users2["name"]["first"]) //Leonardo

	delete(users2, "name")
	users2["signo"] = map[string]string{
		"nome": "gemeos",
	}

}
