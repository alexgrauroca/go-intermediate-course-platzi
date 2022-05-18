package structs

type Employee struct {
	id int
	Human
}

func NewEmployee(id int, name string, lastname1 string, lastname2 string) *Employee {
	e := new(Employee)

	e.SetId(id)
	e.SetName(name)
	e.SetLastname1(lastname1)
	e.SetLastname2(lastname2)

	return e
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}
