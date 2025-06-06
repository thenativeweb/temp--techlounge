package auth

type User struct {
	name  string
	email string
}

func NewUser(name string, email string) *User {
	return &User{
		name:  name,
		email: email,
	}
}

func (user *User) Name() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) Email() string {
	return user.email
}

func (user *User) FullContact() string {
	return user.name + " <" + user.email + ">"
}
