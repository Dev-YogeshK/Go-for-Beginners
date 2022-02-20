// Author : Yogesh Kale
//package main
package main

//Imports
import (
	"fmt"
	"log"
)

//Main Function
func main() {
	fmt.Println("We are Learning Loops")
	// For block
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	fmt.Println("Looping Over Slices")
	// Short Hand For Slices
	mySlices := []string{"Zero", "one", "two", "three", "four", "five"}
	// looping over Slices
	for i, l := range mySlices {
		fmt.Println(i, l)
	}

}
