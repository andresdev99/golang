package main

import (
	"fmt"
	"math/rand"
	. "modgoland/map/utils"
	"strconv"
)

func main() {
	mockData := MockDataStore{User: make(map[int]User)}
	u := User{Id: 0, Name: "John", Age: 30}

	if err := mockData.SaveUser(u); err != nil {
		fmt.Println(err)
	}
	u = User{Id: 1, Name: "Andres", Age: 25}
	if err := mockData.SaveUser(u); err != nil {
		fmt.Println(err)
	}

	for i := len(mockData.User); i <= 20; i++ {
		u = User{Id: i, Name: "Andres" + strconv.Itoa(i), Age: rand.Int()}
		if err := mockData.SaveUser(u); err != nil {
			fmt.Println("There was an error")
		}
	}

	mockData.PrintAllUsers()
}
