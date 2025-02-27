package utils // nama paket

import "fmt" // import paket fmt

func Array() { // fungsi Array
	fmt.Println("")
	fmt.Println("Array")
	fmt.Println("=====")

	// Array dengan ukuran tetap
	names := [3]string{"Vandy", "Rizky", "Pratama"} // inisialisasi array
	fmt.Println(names)                              // print array
	fmt.Println(names[0])                           // print array dengan index 0

	// Array multidimensi
	matrix := [2][2]int{{1, 2}, {3, 4}} // inisialisasi array 2 dimensi
	fmt.Println(matrix)                 // print array 2 dimensi
	fmt.Println(matrix[1][0])           // print array 2 dimensi dengan index 1 dan 0

	// Mengubah Array
	names[0] = "Vandy Rain" // ubah array dengan index 0
	fmt.Println(names)      // print array

	// Menghapus Array
	names[0] = ""      // hapus array dengan index 0
	fmt.Println(names) // print array
}
