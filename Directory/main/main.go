package main

import (
	"fmt"

	"github.com/bayuiqballl/Data-Structure-Golang/Directory/models"
)

func main() {

	user1 := models.Member{
		Name: "Bayu",
		Age:  21,
	}

	fmt.Println("Nama orang ini ", user1)

	//change
	user2 := user1
	user2.Name = "Bayu Muhammad Iqbal"

	//cek
	if user1 == user2 {
		fmt.Println("Sama ")
	} else {
		fmt.Println("Beda, nama orang ini ", user2)
	}

}
