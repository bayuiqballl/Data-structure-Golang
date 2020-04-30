package main

import "fmt"

func main() {

	days := [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	for index, value := range days {
		fmt.Println("Hari ", index, " = ", value)
	}

}
