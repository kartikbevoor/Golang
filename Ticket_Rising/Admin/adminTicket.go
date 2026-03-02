package admin

import (
	"Ticket_Rising/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reply struct {
	TicketId int
	Comment  string
}

var Replies []Reply

func AdminTicket() {
	for {
		fmt.Println("Press 1 to view tickets")
		fmt.Println("Press 2 to reply to ticket")
		fmt.Println("Press 3 to exit")

		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Tickets of your category are:")
			ViewTicketsToAdmin()
		case 2:
			fmt.Println("Select ticket to reply")
			ReplyToTicket()
		case 3:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

//  TicketRiser string
// 	TicketId    int
// 	Category    string
// 	Description string
// 	TicketTime  string
// 	TicketDate  string

func ViewTicketsToAdmin() {
	ThisCategory := getAdminCategory()
	if len(utils.Tickets) == 0 {
		fmt.Println("No tickets to view")
		return
	}
	ticketsViewCount := 0
	for _, v := range utils.Tickets {
		if v.Category == ThisCategory {
			fmt.Println("Ticket Id:", v.TicketId)
			fmt.Println("Date of creation:", v.TicketDate)
			fmt.Println("Time of creation:", v.TicketTime)
			fmt.Println("Description:", v.Description)
			ticketsViewCount++
		}
		if ticketsViewCount == 0 {
			fmt.Println("No tickets of your category")
		}
	}
}

func getAdminCategory() string {
	for _, v := range Admins {
		if ThisAdmin == v.UserName {
			return v.category
		}
	}
	return ""
}

func ReplyToTicket() {
	var tempReply Reply
	fmt.Println("Reply to ticket")

	fmt.Println("Enter ticket id")
	fmt.Scan(tempReply.TicketId)

	fmt.Println("Comment to raised ticket")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give description for your ticket (press ENTER twice to finish):")

	var lines []string
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")

		if line == "" { // stop on empty line
			break
		}
		lines = append(lines, line)
	}

	comment := strings.Join(lines, "\n")
	tempReply.Comment = comment
	// fmt.Scan(&tempReply.comment)

	isvalid := isValidReply(tempReply.TicketId)
	if isvalid {
		Replies = append(Replies, tempReply)
	}
}

func isValidReply(ticketId int) bool {
	for _, v := range utils.Tickets {
		if v.TicketId == ticketId {
			return true
		}
	}
	return false
}
