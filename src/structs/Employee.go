package structs

type Employee struct {
	id        int
	name      string
	lastname1 string
	lastname2 string
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

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) SetLastname1(lastname1 string) {
	e.lastname1 = lastname1
}

func (e *Employee) GetLastname1() string {
	return e.lastname1
}

func (e *Employee) SetLastname2(lastname2 string) {
	e.lastname2 = lastname2
}

func (e *Employee) GetLastname2() string {
	return e.lastname2
}
