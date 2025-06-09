package main

import "fmt"

type User struct {
	Name string
	Age  int
	// X, Y int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

type Users []*User

func main() {

	user1 := new(User)
	fmt.Println(user1)

	user1.Name = "John"
	user1.Age = 20
	fmt.Println(user1)

	user2 := User{Name: "John", Age: 20}
	fmt.Println(user2)

	user3 := NewUser("John", 20)
	fmt.Println(user3)
	fmt.Println(*user3)

	user4 := NewUser("John", 20)
	user5 := NewUser("ttt", 40)
	user6 := NewUser("aaa", 50)

	users := Users{}
	users = append(users, user1, user3, user4, user5, user6)
	fmt.Println(users)

	for _, user := range users {
		fmt.Println(*user)
	}

}
