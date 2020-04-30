package main

import "fmt"

func main() {

	var member [3]string

	member[0] = "Hilman"
	member[1] = "Bayu"
	member[2] = "Juliden"

	daftarAngka := [...]int{10, 20, 30, 40, 50}
	daftarAngka[1] = 100

	fmt.Println(daftarAngka)
	fmt.Println(member)

}
