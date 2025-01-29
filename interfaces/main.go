package main

import (
	"fmt"
	"strconv"
)

type tool struct {
	name string
}

type secretAgent struct {
	person
	tool
	licenseToKill bool
}
type person struct {
	tool
	name string
	age  int
}

type human interface {
	toolsName() string
}

func (p person) toolsName() string {
	return "I'm " + p.name + " my age is " + strconv.Itoa(p.age) + " and my tool is: " + p.tool.name
}

func (p secretAgent) toolsName() string {
	return "I'm a secret agent named " + p.person.name + " my age is " + strconv.Itoa(p.person.age) + " and my tool is: " + p.tool.name
}

func sayTool(pe human) {
	fmt.Println(pe.toolsName())
}

func main() {
	secretAgentObj := secretAgent{
		person: person{
			name: "James Bond",
			age:  32,
		},
		tool: tool{
			name: "Gun",
		},
		licenseToKill: true,
	}

	personObj := person{
		tool: tool{
			name: "Sword",
		},
		name: "Andres",
		age:  25,
	}
	sayTool(secretAgentObj)
	sayTool(personObj)
}
