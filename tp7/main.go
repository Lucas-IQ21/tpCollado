package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	var bob, alice person

	bob = person{"bob", 20}
	bob.age = 18

	alice = person{
		name: "Alice",
		age:  30,
	}

	p := newPerson("Sam")
	q := p
	p.age = 58
	q.age = 60

	fmt.Println(bob, alice, p, q)

}
