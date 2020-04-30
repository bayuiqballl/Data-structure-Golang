package main

import "fmt"

func main() {
	member := [3]string{"Bayu", "Muhammad", "Iqbal"}

	for i := 0; i < len(member); i++ {
		fmt.Println("Membernya " + member[i])
	}

}
