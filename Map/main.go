package main

import "fmt"

// MAP
var member = map[int]string{
	12345: "Bayu",
	56789: "Iqbal",
	10111: "Rahmat",
	12131: "Fauzi",
	14141: "Ganjar",
}

func check(id int) bool {
	_, exists := member[id]
	return exists
}

func main() {

	var newMember = member
	newMember[21212] = "Seto"

	fmt.Println("----------Check Member -----------")
	if check(999999) {
		fmt.Println("member ada")
	} else {
		fmt.Println("member tidak ada")
	}

	// fmt.Println("----------Delete Member -----------")
	delete(member, 14141)

	fmt.Println("----------Tampilan Member -----------")
	for id, name := range member {
		fmt.Println(id, name)
	}

}
