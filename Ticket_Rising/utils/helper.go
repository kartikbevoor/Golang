package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ThisUser string

type Ticket struct {
	TicketRiser string
	TicketId    int
	Category    string
	Description string
	TicketTime  string
	TicketDate  string
}

var Tickets []Ticket

func (t *Ticket) RaiseTicket() Ticket {
	t.TicketRiser = ThisUser

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

func (t *Ticket) ValidatingUserTickets() (bool, []error) {
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

func GenerateUserId() int {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Int())
	n := rand.Intn(9999)
	return n
}
