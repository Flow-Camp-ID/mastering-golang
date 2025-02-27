package utils // nama paket

import "fmt" // import paket fmt

// Pointer
func Pointer() { // fungsi Pointer
	fmt.Println("")
	fmt.Println("Pointer")
	fmt.Println("=======")

	// Pointer
	var numberA int = 4                          // inisialisasi variabel numberA
	var numberB *int = &numberA                  // inisialisasi pointer numberB dengan alamat numberA
	fmt.Println("numberA (value)   :", numberA)  // print numberA
	fmt.Println("numberA (address) :", &numberA) // print alamat numberA
	fmt.Println("numberB (value)   :", *numberB) // print numberB
	fmt.Println("numberB (address) :", numberB)  // print alamat numberB

	// Mengubah Value Pointer
	numberA = 5                                  // ubah value numberA
	fmt.Println("numberA (value)   :", numberA)  // print numberA
	fmt.Println("numberA (address) :", &numberA) // print alamat numberA
	fmt.Println("numberB (value)   :", *numberB) // print numberB
	fmt.Println("numberB (address) :", numberB)  // print alamat numberB
}
