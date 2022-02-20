// Author : Yogesh Kale
//package main
package main

//Imports
import (
	"fmt"
)

//Main Function
func main() {
	fmt.Println("We are Learning loops over map")

	//map of String
	myMap := make(map[string]string)
	// Map Contents
	myMap["zero"] = "Zero 0"
	myMap["one"] = "first 1"
	myMap["two"] = "two 2"
	myMap["three"] = "three 3"
	myMap["four"] = "four 4"

	// Looping Over map Using Blank Identifier
	for _, l := range myMap {
		fmt.Println(l)
	}

	fmt.Println("Looping Over Slices")
	// Short Hand For Slices
	mySlices := []string{"zero", "one", "two", "three", "four", "five"}
	// looping over Slices with blank Identifier
	for _, l := range mySlices {
		fmt.Println(l)
	}

}
