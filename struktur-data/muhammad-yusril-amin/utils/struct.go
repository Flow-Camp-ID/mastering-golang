package utils

import "fmt"

// definisikan struct
type Wisata struct {
	Nama string
	Lokasi string
	Provinsi string
}

// definisikan struct
type Komoditas struct {
	komoditas []string
}

// definisikan struct
type KomoditasWisata struct {
	data Wisata
	komoditasWisata Komoditas
}

func StructFunction(){
	// Membuat instance struct
	detailWisata := KomoditasWisata{
		data : Wisata{
			Nama: "Kawah Ijen",
			Lokasi : "Banyuangi",
			Provinsi : "Kopi",
		},
		komoditasWisata: Komoditas{
			komoditas: []string{"Belerang", "Kopi"},
		},
	}

	// cetak struct
	fmt.Println(detailWisata)
}