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

	fmt.Println(DisplayMe()) // Calling Display Me

	fmt.Println(DisplayInt()) // Calling Display US

}

// function body DisplayMe
func DisplayMe() string {
	return "Hello , I am Display me"
}

// function body DisplayInt
func DisplayInt() int {
	fmt.Println("Hello, I am Display Int with Value Below")
	return 5
}
