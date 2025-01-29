package utils

import (
	"fmt"
)

func PrintPerson(p Greeter) {
	fmt.Println(p.Greet())
}
