// Author : Yogesh Kale
//package main
package main

//Imports
import (
	"fmt"
	"log"
)

// Main Function
func main() {
	// Choice variable for switch
	var choice int
	fmt.Println("Enter the choice :::::")
	fmt.Scanln(&choice) // Accepting Choice

	//Switch Block
	switch choice {
	case 1:
		log.Println("January")
	case 2:
		log.Println("february")
	case 3:
		log.Println("March")
	case 4:
		log.Println("April")
	case 5:
		log.Println("May")
	case 6:
		log.Println("June")
	case 7:
		log.Println("July")
	case 8:
		log.Println("August")
	case 9:
		log.Println("September")
	case 10:
		log.Println("October")
	case 11:
		log.Println("November")
	case 12:
		log.Println("December")
	default: // Default case
		log.Println("Wrong Choice")
	}

}
