// Author : Yogesh Kale
//package main
package main

// imports
import "log"

// Main Function
func main() {

	var myString string // String variable
	myString = "Mystring in main"

	log.Println("myString is set to", myString)                 // Initial String
	changeUsingPointer(&myString)                               // Function Call
	log.Println("after func call myString is set to", myString) // String after Function Call
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "String in changeUsingPointer"
	*s = newValue // reference to newValue
}
