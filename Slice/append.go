package main

import "fmt"

func main() {
	//  Slice
	member := []string{"Bayu", "Muhammad", "Iqbal", "Dani", "Seto"}

	newMember := make([]string, 5)
	copy(newMember, member)

	newMember[1] = "Lukman"
	newMember = append(newMember, "Rina", "Mikha")

	fmt.Println(newMember)
	fmt.Println(member)
}
