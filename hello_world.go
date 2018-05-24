package main

import "fmt"

// SalutationWord is User Defined Type based off of string
type SalutationWord string

// Salutation is an exported struct
type Salutation struct {
	name     string
	greeting string
}

const (
	// PI is exported
	PI = 3.14159
	// Language is exported
	Language = "Go"
)

const (
	// A is exported and an increasing constant
	A = iota
	// B is exported and an increasing constant
	B
	// C is exported and an increasing constant
	C
)

// CreateMessage is exported
func CreateMessage(firstPart string, secondPart string) {
	fmt.Println(firstPart + " " + secondPart)
}

func main() {
	var message = "Hello Go World"
	var greeting *string = &message
	var s = Salutation{"Joe", "Hello"}
	fmt.Println(message, *greeting)
	fmt.Println(s.name)
	fmt.Println(s.greeting)
	fmt.Println(PI, Language)
	fmt.Println(A, B, C)
	// User Defined Types are used instead of classes
	// structs are typically used as UDT

}
