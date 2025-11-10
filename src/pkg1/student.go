package pkg1

import "fmt"

type Student struct {
	User
	School string
}

type Greeter interface {
	Greet() string
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}
