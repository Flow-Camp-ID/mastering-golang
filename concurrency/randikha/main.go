package main

import (
	concurr "concurrency-go/utils"
)

func main() {

	// Memasukkan data orderan pencucian kendaraan
	Orders := []concurr.Order{
		{1, "Toyota Yaris", "B1234UGB", 5},
		{2, "KIA Seltos", "B8123ABH", 10},
		{3, "Hyundai Creta", "B7645GHS", 15},
		{4, "Honda HRV", "B8977AAC", 15},
		{5, "Wuling Almaz", "B9900FCD", 10},
		{6, "Suzuki Ertiga", "B5426FYP", 5},
	}

	concurr.ConcurrencyCase(Orders)
}
