package main

import "fmt"

// Methods
// fungsi, yg menrima dari object tertentu
type Member struct {
	name string
	age  int
}

func (m Member) getInfo() {
	fmt.Println("Halo ini diprint User")
}

func main() {
	user := Member{"Bayu", 21}
	user.getInfo()
	fmt.Println()
}
