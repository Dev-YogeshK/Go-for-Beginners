// Author : Yogesh Kale
//package main
package main

//imports
import (
	"fmt"
)

//Global Variable
var iAmGlobal string = " I am Global Variable"

//Main  function
func main() {

	fmt.Println("We are Learning Functions and Variables in Go")

	fmt.Println(DisplayMe()) // Calling Display Me

}

// function body DisplayMe with two return types
func DisplayMe() (string, int) {
	fmt.Println(iAmGlobal)
	return "Hello , I am Display me", 5
}
