package admin

import "fmt"

func adminStarting() {

}

func OldAdminOrNew() {
	for {
		fmt.Println("Admin Portal")
		fmt.Println("Press 1 to register")
		fmt.Println("Press 2 to login")
		fmt.Println("Press 3 to go back")

		var whichAdmin int
		fmt.Scan(&whichAdmin)

		switch whichAdmin {
		case 1:
			isRegistered := false
			for !isRegistered {
				fmt.Println("Register as new admin")
				var TempAdmin Admin
				TempAdmin = TempAdmin.Register()
				isValidAdmin, err := TempAdmin.validatingAdminRegistration()
				if !isValidAdmin {
					for _, v := range err {
						fmt.Println(v)
					}
				} else {
					Admins = append(Admins, TempAdmin)
					fmt.Println("Registration successful")
					fmt.Println("Login using your login credentials.")
					isRegistered = true
				}
			}
		case 2:
			fmt.Println("Login to your account")
			AdminLogin()
		case 3:
			fmt.Println("Back")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}

}
