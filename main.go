package main

import (
	"fmt"
	// "github.com/LordCeilan/introToGo/packages/convertor"
	// "github.com/LordCeilan/introToGo/packages/inout"
	// "github.com/LordCeilan/introToGo/packages/variables"
	"github.com/LordCeilan/introToGo/packages/pointers"
	// "github.com/LordCeilan/introToGo/packages/variables"
)

func main() {
	printGreeting()
	fmt.Println("Printing Types")
	// variables.MyTypes()
	// convertor.VarConvertor()
	// inout.MyInOut()
	// pointers.NoPointers()
	// pointers.IntPointers()
	// pointers.IntDirection()
	// pointers.DereferencePointers()
	// pointers.WithPointers()
	pointers.StructPointers()
}

func printGreeting() {
	greeting := "Hello there"
	fmt.Println(greeting)
}
