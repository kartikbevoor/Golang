package admin

import (
	"Ticket_Rising/utils"
	"errors"
	"fmt"
)

type Admin struct {
	Name     string
	Id       int
	category string
	UserName string
	password string
}

var Admins []Admin

func (a *Admin) Register() Admin {
	fmt.Println("Enter your name")
	fmt.Scan(&a.Name)
	fmt.Println("Enter your category")
	fmt.Scan(&a.category)
	fmt.Println("Enter your Username")
	fmt.Scan(&a.UserName)
	fmt.Println("Enter your Password")
	fmt.Scan(&a.password)
	// fmt.Println(a)

	a.Id = utils.GenerateUserId()

	return *a
}

func (a *Admin) validatingAdminRegistration() (bool, []error) {
	var err []error
	isValid := true
	if len(a.Name) < 2 {
		isValid = false
		err1 := errors.New("Invalid name: Name should be of atleast 2 character")
		err = append(err, err1)
	}
	if len(a.password) != 4 {
		isValid = false
		err2 := errors.New("Password should be of 4 digits")
		err = append(err, err2)
	}
	if len(a.UserName) < 2 {
		isValid = false
		err3 := errors.New("Invalid username: Userame should be of atleast 2 character")
		err = append(err, err3)
	}
	for _, v := range Admins {
		if v.UserName == a.UserName {
			isValid = false
			err4 := errors.New("Invalid username: Username already exists")
			err = append(err, err4)
		}
	}
	return isValid, err
}
