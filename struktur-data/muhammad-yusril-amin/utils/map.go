package utils

import "fmt"

func MapFunction(){
	// case 1 - map dengan mendefiniskan satu tipe data studi case lokasi
	lokasi := map[string]string {
		"provisi": "Jawa Timur",
		"kabupaten": "Bondowoso",
		"komoditi": "Kopi",
	}
	//  cetak vaiabel
	fmt.Println(lokasi["provisi"])
	
	// case 2 - map dengan multiple tipe data studi case komoditas
	buah := map[string]interface{} {
		"nama" : "Mangga",
		"harga/kg" : 25000,
		"manis" : true,
		"jenis" : []string{"mangga arumanis","mangga apel","mangga golek"},
	}
	fmt.Println(buah)
}