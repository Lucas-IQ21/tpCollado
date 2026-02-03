package main

import "fmt"

func main() {
	if 7%2 == 0 {
		println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	var num int

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num > 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	if num > 0 && num < 10 {
		fmt.Println(num, "is between 1 and 9")
	} else {
		fmt.Println(num, "is another value : ", num)
	}
}
