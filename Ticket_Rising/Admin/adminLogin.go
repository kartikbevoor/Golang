package admin

import "fmt"

var ThisAdmin string

func AdminLogin() {
	isAdminLogin := false
	for !isAdminLogin {
		fmt.Println("Enter your login details")
		var username string
		var password string

		fmt.Println("Enter Username")
		fmt.Scan(&username)

		fmt.Println("Enter Password")
		fmt.Scan(&password)

		isAdminLogin = isValidAdminCredentials(username, password)

		if isAdminLogin {
			ThisAdmin = username
			AdminTicket()
		}

	}
}

func isValidAdminCredentials(username string, password string) bool {
	for _, v := range Admins {
		if v.UserName == username && v.password == password {
			return true
		}
	}
	fmt.Println("Invalid login credentials recheck username and password")
	return false
}
