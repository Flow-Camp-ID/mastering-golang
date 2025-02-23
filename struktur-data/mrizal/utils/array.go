package utils

import "fmt"

func ArrayImplement() {
	// case 1- fixed array
	var arrNumbers [5]int

	arrNumbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arrNumbers)

	// case 2 - array undecided size dengan berbagai tipe data yang berbeda dengan interface{}
	arrNotFix := [...]interface{}{"satu", 2, 3.0, false}
	
	for _, value:=range arrNotFix  {
		fmt.Println(value)
	}

}