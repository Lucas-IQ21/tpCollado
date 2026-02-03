package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	m["k2"] = 13
	fmt.Println("append:", m)
	delete(m, "k1")
	fmt.Println("delete:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
