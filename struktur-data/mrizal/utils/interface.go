package utils

import "fmt"

type Animals interface {
	voice() string 
}

type Anjing struct {
	name string
	age int
}

// struct Anjing yang menerapkan interface dengan fungsi voice() 
func (a Anjing) voice() string {
	return "woof woof"
}

func InterfaceImplement() {
	// interface kosong menerima berbagai data
	arrayInterface := []interface{}{
		"Kalimat atau string",
		30,
		4.5,
		false,
	}

	for _, value := range arrayInterface {
		fmt.Println(value)
	}
	fmt.Println()
	asu := Anjing{name: "Puddy", age: 3}
	fmt.Printf("Anjing bernama %s dengan umur %d suaranya : %s", asu.name, asu.age, asu.voice())
}