package user

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
