package main

import (
	"fmt"
	"sort"
)

func ArrayFunction() {
	fmt.Println("----- Array Function -----")
	// case 1 array fixed size => angka 5 adalah jumlah elemen array
	var numbers [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("ini adalah isi dari numbers: ", numbers)

	//case 2 interface array
	nilaiCampuran := [...]interface{}{1, "string", 3.14, true}
	fmt.Println("ini hasil dari nilai campuran : ", nilaiCampuran)

	//case 3 array spesifik init array
	minumanBotolan := [5]string{"Aqua", "Pocari Sweat", "Teh Botol", "Fanta", "Sprite"}
	fmt.Println("ini hasil dari minumanBotolan ke index 1 : ", minumanBotolan[1])

	//case 4 array multidimensi
	matrix := [2][3]int{{1, 5, 2}, {9, 2, 5}}
	fmt.Println("ini hasil dari matrix: ", matrix[0][2])
}

func MapFunction() {
	namaNegara := map[string]string{
		"ID": "Indonesia",
		"MY": "Malaysia",
		"SG": "Singapura",
	}
	fmt.Println("----- Map Function -----")
	fmt.Println(namaNegara["ID"])

	// case 2 deklarasi map secara literal
	mataUang := map[string]string{
		"ID": "Rupiah",
		"MY": "Ringgit",
		"SG": "Dollar",
	}
	fmt.Println(mataUang)

	// Case 3 map with multiple type data
	makananKesukaan := map[string]interface{}{
		"Jenis":  "Makanan",
		"Nama":   "Nasi Goreng",
		"Harga":  15000,
		"Pedas":  true,
		"Jumlah": 10,
	}

	fmt.Println(makananKesukaan)

	// Case 4 map short ascending
	namaKecamatan := map[string]string{
		"Kec4": "Kecamatan 4",
		"Kec2": "Kecamatan 2",
		"Kec3": "Kecamatan 3",
	}
	keys := make([]string, 0, len(namaKecamatan))
	for k := range namaKecamatan {
		keys = append(keys, k)

	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, namaKecamatan[k])
	}
}

func SliceFunction() {
	// membuat slice
	makanan := [5]string{"Nasi", "Mie", "Soto", "Bakso", "Sate"}
	slice1 := makanan[1:3] // slice dari index 1 sampai 3 hasil [ "Mie", "Soto"]
	slice2 := makanan[:3]  // slice dari index 0 sampai 3 hasil ["Nasi", "Mie", "Soto"]
	fmt.Println("----- Slice Function -----")
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
}

type Sekolah struct {
	KelasA int
	KelasB int
}

func StructFunction() {
	var sekolah Sekolah
	sekolah.KelasA = 15
	sekolah.KelasB = 12
	fmt.Println("----- Struct Function -----")
	fmt.Println("ini adalah jumlah kelas A: ", sekolah.KelasA)
	fmt.Println("ini adalah jumlah kelas B: ", sekolah.KelasB)
}

func Variables() {
	var integerNum = 20

	var floatNum = 20.0

	var complexNum = 1 + 2i
	fmt.Println("----- Variable Function -----")
	fmt.Println(" ini adalah variable number: ", integerNum)
	fmt.Println(" ini adalah variable float: ", floatNum)
	fmt.Println(" ini adalah variable complex: ", complexNum)

}

func main() {
	Variables()
	ArrayFunction()
	MapFunction()
	SliceFunction()
	StructFunction()
}
