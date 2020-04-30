package main

import "fmt"

type Member struct {
	name string
	age  int
}

func main() {

	user1 := Member{
		name: "Bayu",
		age:  21,
	}

	fmt.Println("Nama orang ini ", user1)

	//change
	user2 := user1
	user2.name = "Bayu Muhammad Iqbal"

	//cek
	if user1 == user2 {
		fmt.Println("Sama ")
	} else {
		fmt.Println("Beda, nama orang ini ", user2)
	}

}
