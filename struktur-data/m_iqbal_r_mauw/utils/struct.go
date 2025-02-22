package utils // nama paket

import "fmt" // import paket fmt

// Struct dengan banyak type data
type Person struct {
	Name    string
	Address string
	phone   int
	age     int
	status  bool
} // struct Person

func Struct() { // fungsi Struct
	fmt.Println("")
	fmt.Println("Struct")
	fmt.Println("======")

	// membuat struct
	person := Person{
		Name:    "Vandy",
		Address: "Indonesia",
		phone:   1234567890,
		age:     25,
		status:  true,
	} // inisialisasi struct
	fmt.Println(person) // print struct

	// Mengubah Struct
	person.Name = "Vandy Rain" // ubah struct dengan key Name
	fmt.Println(person)
}
