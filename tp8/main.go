package main

import "fmt"

type User struct {
	Name  string
	Login string
}

func New(name, login string) *User {
	return &User{Name: name, Login: login}
}
func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) GetName() string {
	return u.Name
}

func main() {
	u := New("John Doe", "jdoe@google.com")
	fmt.Println("user: ", u)
	u.UpdateName("Jane Doe")
	fmt.Println("user: ", u.GetName())
}
