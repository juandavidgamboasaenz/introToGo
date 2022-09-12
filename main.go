package main

import (
	"fmt"

	"github.com/LordCeilan/introToGo/types"
)

func main() {
	printGreeting()
	fmt.Println("Printing Types")
	types.MyTypes()
}

func printGreeting() {
	greeting := "Hello there"
	fmt.Println(greeting)
}
