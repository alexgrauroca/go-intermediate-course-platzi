package main

import (
	"fmt"
	"go-intermediate-course-platzi/src/structs"
)

func main() {
	// Returning a pointer to the instance
	e := new(structs.Employee)
	e.SetId(1)
	e.SetName("Àlex")
	e.SetLastname1("Grau")
	e.SetLastname2("Roca")
	fmt.Printf("%+v\n", *e)
}
