// Author : Yogesh Kale
//package main
package main

//imports
import (
	"fmt"
	"sort"
)

// Main Function
func main() {
	fmt.Println("We are learning Slices")

	var mySlices []string // creating Slices
	// appending Values to mySlices
	mySlices = append(mySlices, "Yogesh")
	mySlices = append(mySlices, "Kaushal")
	mySlices = append(mySlices, "Ashwini")

	iAmSlices := []int{5, 4, 8, 9, 1} // Short hand for creating Slices

	sort.Ints(iAmSlices) // Sorted slices to get numbers in ascending Order
	fmt.Println("Values in mySlices", mySlices)
	fmt.Println("Values in iAmSlices", iAmSlices)
	// Printling specific index Values
	fmt.Println("Printing Values of iAmSlices from 0 to 2", iAmSlices[0:2])
}
