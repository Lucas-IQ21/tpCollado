package main

import (
	"fmt"
	"tp12/user"
)

func main() {
	u := user.New("John Doe", "jdoe")
	fmt.Println("User: ", u)
	u.UpdateName("Jane Doe")
	fmt.Println("User: ", u.GetName())
}
