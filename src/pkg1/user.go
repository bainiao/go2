package pkg1

import (
	"cmp"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

func (u User) SetAge(age int) {
	u.Age = age
	switch cmp.Compare(u.Age, 12) {
	case -1:
		fmt.Println("I am a child.")
	case 0:
		fmt.Println("I am 12 years old.")
	case 1:
		fmt.Println("I am not a child anymore.")
	}
}

func (u User) GetAge() int {
	return u.Age
}
