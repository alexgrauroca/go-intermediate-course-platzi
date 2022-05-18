package main

import (
	"fmt"
	"go-intermediate-course-platzi/src/interfaces"
	"go-intermediate-course-platzi/src/structs"
)

func getMessage(p interfaces.PrintInfo) {
	fmt.Println(p.GetMessage())
}

func main() {
	// Returning a pointer to the instance
	fte := structs.NewFullTimeEmployee(1, "Àlex", "Grau", "Roca")
	fmt.Printf("%+v\n", *fte)

	te := structs.NewTemporaryEmployee(1, "Àlex", "Grau", "Roca")
	fmt.Printf("%+v\n", *te)

	getMessage(fte)
	getMessage(te)
}
