package main

import "fmt"

func main() {

	// membuat
	member := make([]string, 5, 10)

	fmt.Println(member)
	//  cek  panjang
	fmt.Println("Length ", len(member))
	//  cek kapasitas
	fmt.Println("Kapasitas ", cap(member))

}
