package utils

import "fmt"

type User struct {
	Id   int
	Name string
	Age  int
}

type MockDataStore struct {
	User map[int]User
}

func (m MockDataStore) GetUser(id int) (User, error) {
	user, ok := m.User[id]

	if !ok {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}

	return user, nil
}

func (m MockDataStore) GetAllUsers() map[int]User {
	return m.User
}

func (m MockDataStore) SaveUser(user User) error {
	v, ok := m.User[user.Id]

	if ok {
		return fmt.Errorf("user with id %d already exists", v.Id)
	}
	m.User[user.Id] = user
	return nil
}

func (m MockDataStore) PrintAllUsers() {
	users := m.GetAllUsers()

	for i := 0; i < len(users); i++ {
		fmt.Printf("INDEX:%d  ID: %d - Name:%v\n", i, users[i].Id, users[i].Name)
	}
}
