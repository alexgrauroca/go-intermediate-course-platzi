package main

import (
	"fmt"
	"go-intermediate-course-platzi/src/structs"
)

func main() {
	e := structs.Employee{}
	e.SetId(1)
	e.SetName("Àlex")
	e.SetLastname1("Grau")
	e.SetLastname2("Roca")
	fmt.Printf("%+v\n", e)
}
