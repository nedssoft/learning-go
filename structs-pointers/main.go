package main

import (
	"fmt"
)

type nums []int
func main() {
	// n := nums{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// n.EvenOrOdd()
	jimmy := person{
		firstName: "Jimmy",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jimmy@doe.com",
			zipCode: 12345,
		},
	}
	jimmy.updateName("Jimmy Jr.")
	jimmy.print()
}

func (n nums) EvenOrOdd() {
	for _, v := range n {
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}