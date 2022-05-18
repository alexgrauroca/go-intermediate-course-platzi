package structs

type TemporaryEmployee struct {
	Employee
	taxRate int
}

func (te TemporaryEmployee) GetMessage() string {
	return "I'm a temporary employee"
}

func NewTemporaryEmployee(id int, name string, lastname1 string, lastname2 string) *TemporaryEmployee {
	e := new(TemporaryEmployee)

	e.SetId(id)
	e.SetName(name)
	e.SetLastname1(lastname1)
	e.SetLastname2(lastname2)

	return e
}
