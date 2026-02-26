package user

import (
	//"Ticket_Rising/utils"

	"errors"
	"fmt"
	"math/rand"
	"time"
)

var Tickets []Ticket

// var thisUser string

func raiseTicketOrViewTicket() {
	for {
		var raiseOrView int

		fmt.Println("Press 1 to raise ticket")
		fmt.Println("Press 2 to view ticket")
		fmt.Println("Press 3 to logout")

		fmt.Scan(&raiseOrView)

		switch raiseOrView {
		case 1:
			fmt.Println("Raise Ticket")
			isTicketRaised := false
			for !isTicketRaised {
				var tempTicket Ticket
				tempTicket = tempTicket.raiseTicket()
				isValidTicket, err := tempTicket.validatingUserTickets()
				if !isValidTicket {
					for _, v := range err {
						fmt.Println(v)
					}
				} else {
					isTicketRaised = true
					Tickets = append(Tickets, tempTicket)
					fmt.Println("Ticket raised successfully")
				}
			}
			//raiseTicketOrViewTicket()
		case 2:
			fmt.Println("Ticket Raised are")
			for _, v := range Tickets {
				if v.ticketRiser == ThisUser {
					fmt.Println(v)
				}
			}
		case 3:
			fmt.Println("Logout")
			UserOrAdmin()
		}
	}

}

type Ticket struct {
	ticketRiser string
	TicketId    int
	Category    string
	Description string
	TicketTime  string
	TicketDate  string
}

func (t *Ticket) raiseTicket() Ticket {
	t.ticketRiser = ThisUser

	t.TicketId = generateTicketId()

	fmt.Println("Enter your category")
	fmt.Scan(&t.Category)

	fmt.Println("Give description for your ticket")
	fmt.Scan(&t.Description)

	//now := time.Now()
	formatted := time.Now().Format("15:04:05")
	t.TicketTime = formatted

	today := time.Now().Format("2006-01-02")
	t.TicketDate = today

	return *t
}

func generateTicketId() int {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Int())
	n := rand.Intn(9999)
	return n
}

func (t *Ticket) validatingUserTickets() (bool, []error) {
	isValid := true
	var err []error
	if len(t.Category) < 2 {
		err1 := errors.New("Invalid category: to small category")
		isValid = false
		err = append(err, err1)
	}

	if len(t.Description) < 3 {
		err2 := errors.New("Invalid description: to small description")
		isValid = false
		err = append(err, err2)
	}

	return isValid, err
}
