package main

import (
	"fmt"
	"time"
	"sync"
)


const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)


// similiar to Java we can compare struct with class
type UserData struct {
	firstName string
	lastName string
	userEmail string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}


func main() {

	helloworld()

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket(userTickets, firstName, lastName, userEmail)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, userEmail)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n",firstNames)

			if remainingTickets == 0 {
				fmt.Println("OUR CONFERENTE IS BOOKED OUT. COME BACK NEXT YEAR.")
				//break
			}
		} else {
			if !isValidName{
				fmt.Println("FIRST NAME OR LAST NAME YOU ENTERED IS TOO SHORT.")
			}
			if !isValidEmail{
				fmt.Println("EMAIL ADDRESS YOU ENTERED DOESN'T CONTAINS @ SIGN.")
			}
			if !isValidTicketNumber{
				fmt.Println("NUMBER OF TICKETS YOU ENTERED IS INVALID.")
			}
		}
	}
	wg.Wait()
}


func greetUsers(){	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	// when we not have a size of array to run a for loop we can insert a placeholder '_'
	for _, booking := range bookings {		
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&userEmail)

	fmt.Println("Enter your number tickets: ")
	fmt.Scan(&userTickets)

	// In Golang we can return multiple values from function
	return firstName, lastName, userEmail, userTickets
}

func bookTicket (userTickets uint, firstName string, lastName string, userEmail string) {//(string, string, string, uint, uint){
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// maps can not mix data types!
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		userEmail: userEmail,
		numberOfTickets: userTickets,
	}

	//userData["firstNameKey"] = firstName
	//userData["lastNameKey"] = lastName
	//userData["userEmailKey"] = userEmail
	//// we need to convert uint to string
	//userData["numberOfTicketsKey"]	= strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("THANK YOU %v %v FOR BOOKING %v TICKETS. YOU WILL RECEIVE A CONFIRMATION EMAIL AT %v.\n",firstName,lastName,userTickets,userEmail)
	fmt.Printf("%v TICKETS REMAINING FOR %v\n",remainingTickets, conferenceName)

	//return firstName, lastName, userEmail, userTickets, remainingTickets
}


func sendTicket(userTickets uint ,firstName string,lastName string,userEmail string){
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets,firstName,lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, userEmail)
	fmt.Println("#################")
	wg.Done()
}