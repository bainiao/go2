package pkg1

import "fmt"

const (
	Num = iota * 2
	Num2
	Num3
)

type Season uint8

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	default:
		return "unknown"
	}
}

func Printpkg1() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := make([]int, 5, 10)
	mapping := make(map[int]int)
	ch1 := make(chan int, 5)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		slice[i] = arr[i]
		mapping[i] = arr[i]
		ch1 <- arr[i]
	}
	fmt.Println("This is pkg1")
}
