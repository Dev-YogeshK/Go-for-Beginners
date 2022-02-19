// Author : Yogesh Kale
//package main
package main

//imports
import (
	"fmt"
)

//Main  function

func main() {

	fmt.Println("We are Learning Functions in Go")

	DisplayMe() // Calling Display Me

	DisplayUs() // Calling Display US

}

// function body DisplayMe
func DisplayMe() {
	fmt.Println("Hey , I am Display Me")
}

// function body DisplayUs
func DisplayUs() {
	fmt.Println("Hey, I am Display US")
}
