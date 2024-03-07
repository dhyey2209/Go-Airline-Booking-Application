package main

import (
	"airline-booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Global variables and constants for the application
var flightName string = "GoAir Flight 101"
var remainingSeats uint = 120
var bookings = make([]PassengerData, 0)

const totalSeats int = 120

// PassengerData struct to hold booking information.
type PassengerData struct {
	firstName       string
	lastName        string
	email           string
	passportNumber  string
	numberOfSeats   uint
}

// WaitGroup to manage concurrency for sending boarding passes.
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, passportNumber, userSeats := getUserInput()
		isValidName, isValidEmail, isValidPassportNumber, isValidSeatNumber := helper.ValidateUserInput(firstName, lastName, email, passportNumber, userSeats, remainingSeats)

		if isValidName && isValidEmail && isValidPassportNumber && isValidSeatNumber {
			bookSeats(userSeats, firstName, lastName, email, passportNumber)
			wg.Add(1)
			go sendBoardingPass(userSeats, firstName, lastName, email)

			if remainingSeats == 0 {
				fmt.Println("This flight is now fully booked. Please check other flights.")
				break
			}
		} else {
			reportInputErrors(isValidName, isValidEmail, isValidPassportNumber, isValidSeatNumber)
		}
		wg.Wait()
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", flightName)
	fmt.Printf("We have a total of %v seats and %v seats are still available for booking\n", totalSeats, remainingSeats)
	fmt.Println("Book your seats here to fly with us.")
}

func getUserInput() (string, string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var passportNumber string
	var userSeats uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your passport number:")
	fmt.Scan(&passportNumber)

	fmt.Println("Enter the number of seats you want to book:")
	fmt.Scan(&userSeats)

	return firstName, lastName, email, passportNumber, userSeats
}

func bookSeats(userSeats uint, firstName string, lastName string, email string, passportNumber string) {
	remainingSeats -= userSeats

	var passenger = PassengerData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		passportNumber: passportNumber,
		numberOfSeats:  userSeats,
	}

	bookings = append(bookings, passenger)
	fmt.Printf("List of bookings is now: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v seats. You will receive a confirmation email at %v\n", firstName, lastName, userSeats, email)
	fmt.Printf("%v seats remaining for %v\n", remainingSeats, flightName)
}

func sendBoardingPass(userSeats uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	fmt.Println("###################")
	var pass = fmt.Sprintf("%v seats for %v %v", userSeats, firstName, lastName)
	fmt.Printf("Sending boarding pass:\n%v \nto email address %v\n", pass, email)
	fmt.Println("###################")

	wg.Done()
}

func reportInputErrors(isValidName, isValidEmail, isValidPassportNumber, isValidSeatNumber bool) {
	if !isValidName {
		fmt.Println("Your name is invalid, please enter a valid first and last name.")
	}
	if !isValidEmail {
		fmt.Println("Your email is invalid, please include an '@' sign.")
	}
	if !isValidPassportNumber {
		fmt.Println("Your passport number is invalid, please try again.")
	}
	if !isValidSeatNumber {
		fmt.Println("The number of seats you entered is not available, please try again.")
	}
}
