package user

import (
	//"Ticket_Rising/utils"

	"Ticket_Rising/utils"
	"fmt"
)

// var thisUser string

func raiseTicketOrViewTicket() {
	for {
		var raiseOrView int

		fmt.Println("Press 1 to raise ticket")
		fmt.Println("Press 2 to view ticket")
		fmt.Println("Press 3 to view Comments")
		fmt.Println("Press 4 to logout")

		fmt.Scan(&raiseOrView)

		switch raiseOrView {
		case 1:
			fmt.Println("Raise Ticket")
			isTicketRaised := false
			for !isTicketRaised {
				var tempTicket utils.Ticket
				tempTicket = tempTicket.RaiseTicket()
				isValidTicket, err := tempTicket.ValidatingUserTickets()
				if !isValidTicket {
					for _, v := range err {
						fmt.Println(v)
					}
				} else {
					isTicketRaised = true
					utils.Tickets = append(utils.Tickets, tempTicket)
					fmt.Println("Ticket raised successfully")
				}
			}
			//raiseTicketOrViewTicket()
		case 2:
			fmt.Println("Ticket Raised are")
			for _, v := range utils.Tickets {
				if v.TicketRiser == utils.ThisUser {
					fmt.Println(v)
				}
			}
		case 3:
			fmt.Println("Comments")
		case 4:
			fmt.Println("Logout")
			return //UserOrAdmin()
		}
	}

}
