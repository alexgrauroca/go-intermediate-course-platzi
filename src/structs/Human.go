package structs

type Human struct {
	name      string
	lastname1 string
	lastname2 string
	age       int
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) SetLastname1(lastname1 string) {
	h.lastname1 = lastname1
}

func (h *Human) GetLastname1() string {
	return h.lastname1
}

func (h *Human) SetLastname2(lastname2 string) {
	h.lastname2 = lastname2
}

func (h *Human) GetLastname2() string {
	return h.lastname2
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func (h *Human) GetAge() int {
	return h.age
}
