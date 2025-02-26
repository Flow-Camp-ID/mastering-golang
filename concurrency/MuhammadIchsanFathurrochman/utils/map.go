package utils

import (
	"fmt"
	"sort"
)

func MapFunction() {
		// case 1 - deklarasi Map
		bahasaPemrograman := make(map[string]string)

		bahasaPemrograman["Tech"] = "Technology"
		bahasaPemrograman["GORM"] = "Golang ORM"

		fmt.Println(bahasaPemrograman["Tech"])

		// case 2 - deklasi map secara literal
		bootcamp := map[string]string {
			"JS": "JavaScript",
			"PHP": "Bahasa PHP",
			"GO": "Golang",
		}

		fmt.Println(bootcamp["JS"])

		// case 3 - map with multiple type data
		dataMaps := map[string]interface{} {
			"nama": "Taufik",
			"Hobi": []string{"Coding", "Computer"},
			"Role": "Backend",
			"Status": true,
			"Jumlah saudara": 4,
		}

		// langkah 1
		keys := make([]string, 0, len(dataMaps))
		for key := range dataMaps {
			keys = append(keys, key)
		}

		// langkah 2
		sort.Strings(keys)

		// langkah 3
		fmt.Println("\nHasil pengurutan:")
		for _, key := range keys {
			fmt.Printf("Key: %s, Value: %v", key, dataMaps[key])
		}

		// for _, value := range dataMaps {
		// 	fmt.Println(value)
		// }

		// // sebelum
		// fmt.Println(dataMaps)
		// fmt.Println(dataMaps["nama"])

		// delete(dataMaps, "nama")

		// // sesudah
		// fmt.Println(dataMaps)
		// fmt.Println(dataMaps["nama"])

}