// Author : Yogesh Kale
//package main
package main

//Imports
import (
	"fmt"
)

// interface declaration
type Cars interface {
	features() string
	color() string
}

// struct rollroyce
type rollroyce struct {
	name     string
	noOfTyre int
}

// Struct Audi
type Audi struct {
	name     string
	noOfTyre int
	price    float64
}

//Main Function
func main() {
	fmt.Println("We are learning Interface")
	// var of type struct rollroyce
	r := rollroyce{
		name:     "RollsRoyce",
		noOfTyre: 6,
	}
	// var of type Audi
	a := Audi{
		name:     " Audi",
		noOfTyre: 4,
		price:    250000,
	}
	// Displaying Values
	PrintInfo(&r)
	fmt.Println(r)
	PrintInfo(&a)
	fmt.Println(a)

}

// function PrintInfo with arguement passed type Cars i.e  Interface
func PrintInfo(c Cars) {
	fmt.Println("I am a Car having color ", c.color(), " features ", c.features(), " Thank You")
}

// function color with reciever
func (e *rollroyce) color() string {
	return "Yellow"
}

// function features with reciever
func (w *rollroyce) features() string {
	return "Hello Features"
}

// function color with reciever
func (e *Audi) color() string {
	return "Black"
}

// function features with reciever
func (e *Audi) features() string {
	return "Hello Features in Audi"
}
