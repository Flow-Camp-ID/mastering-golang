package utils // nama paket

import "fmt" // import paket fmt

func Map() { // fungsi Map
	fmt.Println("")
	fmt.Println("Map")
	fmt.Println("===")

	// Map
	person := map[string]string{
		"name":    "Vandy",
		"address": "Indonesia",
		"email":   "vandy@email.com",
	} // inisialisasi map
	fmt.Println(person) // print map

	// Print Map Value
	fmt.Println(person["name"]) // print map dengan key name
	for _, biodata := range person {
		fmt.Println(biodata)
	} // print map

	// Mengubah Map
	person["name"] = "Vandy Rain" // ubah map dengan key name
	fmt.Println(person["name"])   // print map dengan key name

	// Menghapus Map
	delete(person, "address") // hapus map dengan key address
	for _, biodata := range person {
		fmt.Println(biodata)
	} // print map
}
