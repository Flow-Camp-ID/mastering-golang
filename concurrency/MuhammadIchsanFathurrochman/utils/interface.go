package utils

import "fmt"

type Shape interface {
	Area() float64
}

type Segitiga struct {
	Lebar float64
	Tinggi float64
}

func (r Segitiga) Area() float64 {
	return r.Lebar * r.Tinggi * 0.5
}

func printShapeInfo(s Shape) {
	fmt.Println("Datanya adalah: ", s.Area())
}

func Interfaces() {
	// case 1 - print segitiga
	rect := Segitiga{Lebar: 10, Tinggi: 20}
	fmt.Println(rect)

	// case 2 menggunakan shape interfaces
	shapes := []Shape{rect}
	for i, shape := range shapes {
		fmt.Println(i)
		printShapeInfo(shape)
	}
}