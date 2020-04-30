package main

import "fmt"

func main() {
	//  Slice || Array
	member := [5]string{"Bayu", "Muhammad", "Iqbal", "Dani", "Seto"}
	// membuat slice dari array yg sudah ada
	// mengakses
	slicePertama := member[:1]
	fmt.Println(slicePertama)
	// merubah isi slice
	sliceLengkap := member[:]
	sliceLengkap[1] = "Angel"

	fmt.Println(member)
	fmt.Println(sliceLengkap)

}
