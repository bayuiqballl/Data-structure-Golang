package main

import "fmt"

// Methods
// fungsi, yg menrima dari object tertentu
type Member struct {
	name string
	age  int
}

func (member *Member) changeInfo(newName string, newAge int) {
	member.name = newName
	member.age = newAge
}

func main() {
	user := Member{"Bayu", 21}
	user.changeInfo("Jokowi", 70)
	fmt.Println(user)
}
