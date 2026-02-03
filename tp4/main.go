package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "Mon"
	s = append(s, "petit")
	s[1] = "Chien"
	fmt.Println("apd:", s)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i * j
		}
	}

	fmt.Println("2d: ", twoD)
}
