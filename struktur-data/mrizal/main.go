package main

import (
	"fmt"
	"mrizal/utils"
)

func main() {

	
	fmt.Println("Array Implementation")
	// jika mau memanggil fungsi dari file yang lain di folder yang sama
	// harus dipanggil bersamaan dengan file main.go jadi misal ingin memanggil fungsi dibawah
	// harus go run main.go array.go
	// ArrayImplement()

	// case 1 - var array fixed
	utils.ArrayImplement()

	fmt.Println("\nSlice Impmenetation")

	utils.SliceImplement()

	fmt.Println("\nMap Implementation")

	utils.MapImplement()

	fmt.Println("\nStruct Implementation")

	utils.StructImplement()

	fmt.Println("\nInterface Implementation")

	utils.InterfaceImplement()
	
}