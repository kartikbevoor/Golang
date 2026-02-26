package main

import (
	//"Ticket_Rising/utils"
	user "Ticket_Rising/User"
	"fmt"
)

var IsLogin bool

func main() {
	IsLogin = false // no one is logged
	fmt.Println("Welcome to Ticket Rising Platform")

	if !IsLogin {
		UserOrAdmin() // no one is logged check who want to login user or admin
	}

}

type CommonLogin interface { // common user and admin login interface
	Login(string, int)
}

type commonLogout interface { // common user and admin logout interface
	Logout()
}

type CommonRegistration interface { // common registration for new user or admin
	Register() struct{}
}

type userRegistrationValidation interface { // common registration details validation
	validatingUserRegistration() (bool, string) // for both user and admin
}

func UserOrAdmin() {
	fmt.Println("Press 1 for user")
	fmt.Println("Press 2 for Admin")
	fmt.Println("Press 3 to exit")

	var who int
	fmt.Scan(&who)

	switch who {
	case 1:
		fmt.Println("User")
		user.UserStarting()
	case 2:
		fmt.Println("Admin")
	case 3:
		fmt.Println("Exit")
		return
	}
}
