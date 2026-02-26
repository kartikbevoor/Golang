package user

import (
	"fmt"
)

var ThisUser string

func UserLogin() {
	isUserLogin := false
	for !isUserLogin {
		fmt.Println("Enter your login details")
		var username string
		var password string

		fmt.Println("Enter Username")
		fmt.Scan(&username)

		fmt.Println("Enter Password")
		fmt.Scan(&password)

		isUserLogin = isValidUserCredentials(username, password)

		if isUserLogin {
			ThisUser = username
			raiseTicketOrViewTicket()
		}
	}
}

func isValidUserCredentials(username string, password string) bool {
	for _, v := range Users {
		if v.UserName == username && v.password == password {
			return true
		}
	}
	fmt.Println("Invalid login credentials recheck username and password")
	return false
}
