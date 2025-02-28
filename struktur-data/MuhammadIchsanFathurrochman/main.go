package main

import "fmt"

type Manusia struct {
	nama         string
	jenisKelamin string
}

// Penggunaan ARRAY LIST
func arrayList() {
	// Digunakan untuk meng-identify nama bulan by index karna jumlah bulan dalam 1 tahun sudah pasti / konstan
	var arrayBulan [12]string = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"}
	fmt.Println(arrayBulan)
	fmt.Println()
}

func sliceList(data string) {
	sliceList := []string{"Muhammad", "Ichsan"}

	sliceList = append(sliceList, data)
	fmt.Println("Berhasil menambahkan data kedalam slice")
	fmt.Println(sliceList)
}

func mapList() {
	wargaBogor := make(map[string]*Manusia)

	wargaBogor["ichsan"] = &Manusia{"Ichsan", "L"}
	wargaBogor["tasya"] = &Manusia{"Tasya", "P"}

	for key, item := range wargaBogor {
		fmt.Printf("%s: %+v\n", key, item)
	}
}

func main() {
	// EXAMPLE ARRAY
	arrayList()

	// EXAMPLE SLICE
	sliceList("Fathurrochman")

	mapList()
}
