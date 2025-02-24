package utils // nama paket

import "fmt" // import paket fmt

// Interface
type HasName interface { // deklarasi interface HasName
	GetName() string // deklarasi method GetName
}

type Human struct { // deklarasi struct Person
	Name string // deklarasi variabel Name
}

func (human Human) GetName() string { // deklarasi method GetName
	return human.Name // return variabel Name
}

func Interface() { // fungsi Interface
	fmt.Println("")
	fmt.Println("Interface")
	fmt.Println("=========")

	// Interface
	var vandy Human              // inisialisasi struct Human
	vandy.Name = "Vandy"         // inisialisasi variabel Name
	fmt.Println(vandy.GetName()) // print vandy.GetName()
}
