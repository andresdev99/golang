package main

import . "modgoland/inheritance/utils"

func main() {
	obj := SecretAgent{
		Person: Person{
			Name: "James Bond",
			Age:  32,
		},
		Tools: Tools{
			Name: "Gun",
		},
		LicenseToKill: true,
	}
	SayTool(obj)
}
