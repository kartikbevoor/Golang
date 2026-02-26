package user

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// slice of users
var Users []User

type User struct {
	Name     string
	Id       int
	UserName string
	password string
}

func (u *User) Register() User {
	fmt.Println("Enter your name")
	fmt.Scan(&u.Name)
	fmt.Println("Enter your Username")
	fmt.Scan(&u.UserName)
	fmt.Println("Enter your Password")
	fmt.Scan(&u.password)
	fmt.Println(u)

	u.Id = generateUserId()

	return *u
}

func generateUserId() int {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Int())
	n := rand.Intn(9999)
	return n
}

func (u *User) validatingUserRegistration() (bool, []error) {
	var err []error
	isValid := true
	if len(u.Name) < 2 {
		isValid = false
		err1 := errors.New("Invalid name: Name should be of atleast 2 character")
		err = append(err, err1)
	}
	if len(u.password) != 4 {
		isValid = false
		err2 := errors.New("Password should be of 4 digits")
		err = append(err, err2)
	}
	if len(u.UserName) < 2 {
		isValid = false
		err3 := errors.New("Invalid username: Userame should be of atleast 2 character")
		err = append(err, err3)
	}
	for _, v := range Users {
		if v.UserName == u.UserName {
			isValid = false
			err4 := errors.New("Invalid username: Username already exists")
			err = append(err, err4)
		}
	}
	return isValid, err
}
