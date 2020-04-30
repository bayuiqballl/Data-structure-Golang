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
	for i := 0; i < len(hewan); i++ {
		for j := 0; j < len(hewan[i]); j++ {
			fmt.Println(hewan[i][j])
		}
		fmt.Println("________")
	}

}
