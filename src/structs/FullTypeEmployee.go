package structs

type FullTimeEmployee struct {
	Employee
}

func (fte FullTimeEmployee) GetMessage() string {
	return "It's a full time employee"
}

func NewFullTimeEmployee(id int, name string, lastname1 string, lastname2 string) *FullTimeEmployee {
	e := new(FullTimeEmployee)

	e.SetId(id)
	e.SetName(name)
	e.SetLastname1(lastname1)
	e.SetLastname2(lastname2)

	return e
}
