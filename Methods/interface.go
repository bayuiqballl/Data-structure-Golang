package main

import (
	"fmt"
	"math"
)

// interface
type Bentuk interface {
	keliling() float64
	luas() float64
}

//  struct
type kotak struct {
	panjang, lebar float64
}

type lingkaran struct {
	radius float64
}

type segitiga struct {
	alas, tinggi float64
}

// method
func (k kotak) keliling() float64 {
	return 2*k.panjang + 2*k.lebar
}

func (k kotak) luas() float64 {
	return k.panjang * k.lebar
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.radius
}

func (l lingkaran) luas() float64 {
	return math.Pi * l.radius * l.radius
}

func (s segitiga) keliling() float64 {
	pangkatA := math.Pow(s.alas, 2)
	pangkatB := math.Pow(s.tinggi, 2)

	return math.Sqrt(pangkatA + pangkatB)

}

func (s segitiga) luas() float64 {
	return 0.5 * s.alas * s.tinggi
}

// function
func hitungBangunan(b Bentuk) {

	fmt.Println("Keliling ", b.keliling())
	fmt.Println("Luas ", b.luas())
}

func main() {
	fmt.Println("---persegi---")
	kotak1 := kotak{5, 10}
	hitungBangunan(kotak1)

	fmt.Println("---lingkaran---")
	lingkaran1 := lingkaran{10}
	hitungBangunan(lingkaran1)

	fmt.Println("---segitiga---")
	segitiga1 := segitiga{5, 10}
	hitungBangunan(segitiga1)

}
