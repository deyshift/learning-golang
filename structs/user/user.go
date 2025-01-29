package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// This is a struct that embeds the User struct
type Admin struct {
	User     // anonymous field allows us to access the fields and methods of the User struct
	email    string
	password string
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

// This is a constructor function for the User struct
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName, and birthdate are required")
	}
	// This is a pointer to a User struct
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// This function takes a user struct by value and outputs the user details
func (user User) OutputUserDetails() {
	// pointers are automatically dereferenced in Go when you access a field of a struct
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

// This function takes a pointer to a User struct and clears the first and last name
func (user *User) ClearUserName() {
	// ...
	user.firstName = ""
	user.lastName = ""
}
