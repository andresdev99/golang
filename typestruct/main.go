package main

import . "modgoland/typestruct/utils"

func main() {
	var per PersonActions = &Person{Name: "John"}
	per.CreatePerson("John")
	PrintPerson(per)
}
