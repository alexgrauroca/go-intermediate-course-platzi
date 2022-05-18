package main

import (
	"fmt"
	"go-intermediate-course-platzi/src/structs"
)

func main() {
	// Returning a pointer to the instance
	e := structs.NewEmployee(1, "Àlex", "Grau", "Roca")
	fmt.Printf("%+v\n", *e)
}
