// Author : Yogesh Kale
//package main
package main

//Imports
import (
	"log"
)

// Struct Declaration
type Profile struct {
	name string
	runs int
	avg  float64
	team string
}

// Main Function
func main() {
	//  user Variable of Type Struct
	user := Profile{
		name: "Yogesh",
		runs: 2000,
		avg:  45.85,
		team: "Lucknow Super Giant",
	}

	// Displaying Values
	log.Println("Name is ", user.name)
	log.Println("Runs Scored", user.runs)
	log.Println("Average is", user.avg)
	log.Println("Team  name is", user.team)

}
