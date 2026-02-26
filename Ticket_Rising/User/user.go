package user

import (
	"fmt"
)

func UserStarting() {
	OldUserOrNew()
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
		UserStarting()
	case 2:
		fmt.Println("Admin")
	case 3:
		fmt.Println("Exit")
		return
	}
}

func OldUserOrNew() {
	fmt.Println("Press 1 to register")
	fmt.Println("Press 2 to login")

	var whichUser int
	fmt.Scan(&whichUser)

	switch whichUser {
	case 1:
		isRegistered := false
		for !isRegistered {
			fmt.Println("Register")
			var TempUser User
			TempUser = TempUser.Register() // TempUser =
			isValidUser, err := TempUser.validatingUserRegistration()
			if isValidUser {
				Users = append(Users, TempUser)
				fmt.Println("Registration successful")
				fmt.Println("Login using your login credentials.")
				isRegistered = true
			} else {
				for _, v := range err {
					fmt.Println(v)
				}
			}
			if isRegistered {
				OldUserOrNew()
			}
		}
	case 2:
		fmt.Println("Login")
		UserLogin()
	default:
		fmt.Println("Invalid Choice")
		OldUserOrNew()
	}

}
