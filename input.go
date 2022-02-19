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

	var name string // Variable Declaration
	fmt.Println("Enter your Name:-")
	fmt.Scanln(&name) // Accepting Input
	var age int
	fmt.Println("Enter your Age:-")
	fmt.Scanln(&age)

	// Displaying  Accepted Values
	log.Println("Name is :-", name)
	log.Println("Age is :-", age)
}
