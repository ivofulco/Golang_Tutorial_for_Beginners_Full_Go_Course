package main

import "strings"

var MyVar = "somevalue"

func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) { 
	//var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
