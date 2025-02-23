package utils

import "fmt"

func SliceImplement() {
	// inisialisasi slice dengan panjang sebesar 5
	myPizza := make([]float32, 5)
	// kesimpulan slice merupakan tipe array yg akan digunakan sebagai deklarasi awal (dinamis size)
	for index := range myPizza {
		myPizza[index] = float32(index) + 0.69
	}

	myPizza = append(myPizza, 2.5)
	myPizza = append(myPizza, 3.3)
	myPizza = append(myPizza, 4.5)
	myPizza = append(myPizza, 5.1)
	myPizza = append(myPizza, 6.2)
	myPizza = append(myPizza, 7.9)
	myPizza = append(myPizza, 8.5)

	fmt.Println(myPizza)

	//  berikut slice of array
	fmt.Println("Menampilkan konten awal slice")
	// menampilkan data dari index ke - 0 sampai dengan index ke - 4
	fmt.Println(myPizza[0:5])
}