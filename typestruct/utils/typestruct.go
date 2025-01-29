package utils

type Greeter interface {
	Greet() string
}

type PersonActions interface {
	Greeter
	CreatePerson(name string)
}

type Person struct {
	Name string
}

func (p *Person) CreatePerson(name string) {
	p.Name = name
}

func (p *Person) Greet() string {
	return "Hello, " + p.Name
}
