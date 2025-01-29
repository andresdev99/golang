package utils

import "fmt"

type Tools struct {
	Name string
}

type SecretAgent struct {
	Person
	Tools
	LicenseToKill bool
}
type Person struct {
	Name string
	Age  int
}

type human interface {
	toolsName() string
}

func (p Person) toolsName() string {
	return "My Tools Is: " + p.Name
}

func (p SecretAgent) toolsName() string {
	return "My Secret Tools Is: " + p.Tools.Name
}

func SayTool(pe human) {
	fmt.Println(pe.toolsName())
}
