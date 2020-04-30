package main

import "fmt"

func main() {
	hewan := [4][2]string{
		{"macan", "singa"},
		{"kucing", "kambing"},
		{"rusa", "babi"},
		{"elang", "merpati"},
	}

	// [00][01]
	// [10][11]
	// [20][21]

	fmt.Println(hewan)
	fmt.Println(hewan[1][1])

}
