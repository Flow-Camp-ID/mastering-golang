package utils // nama paket

import "fmt" // import paket fmt

func Slice() { // fungsi Slice
	fmt.Println("")
	fmt.Println("Slice")
	fmt.Println("=====")

	// Slice dengan ukuran dinamis
	names := []string{"Vandy", "Rizky", "Pratama"} // inisialisasi slice
	fmt.Println(names)                             // print slice

	// Print Range Slice
	fmt.Println(names[0:2]) // print slice dengan range 0 sampai 2
	fmt.Println(names[1:3]) // print slice dengan range 1 sampai 3
}
