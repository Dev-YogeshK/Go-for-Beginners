// Author : Yogesh Kale
//package main
package main

//imports
import "log"

// Struct
type Persons struct {
	name  string
	age   int
	place string
}

// Main Function
func main() {
	iAmMap := make(map[string]string)     // map of type String
	iAmMapInt := make(map[string]int)     // map of type int
	iAmMapStruct := make(map[int]Persons) // map of type struct

	// Assigning Value to type Struct
	st := Persons{
		name:  "Rahul",
		age:   34,
		place: " Nashik",
	}

	iAmMap["Yogesh"] = "I am key holding name Yogesh with type String"
	iAmMapInt["one"] = 1
	iAmMapStruct[2] = st

	// Displaying Values
	log.Println(iAmMap["Yogesh"])
	log.Println("I am number of type int and value :- ", iAmMapInt["one"])
	log.Println(" Values stored in map type Struct:- ")
	log.Println(iAmMapStruct[2].name, iAmMapStruct[2].age, iAmMapStruct[2].place)

}
