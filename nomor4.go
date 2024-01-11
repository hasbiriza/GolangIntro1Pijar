package main

import (
	"fmt"
)

func cekKondisi(tinggiSegitiga int) int {
	for i := tinggiSegitiga; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Print("\n")
	}
	return tinggiSegitiga
}

func main() {
	//validasi
	tinggiSegitiga := 5
	if tinggiSegitiga <= 0 {
		fmt.Println("Tinggi segitiga harus lebih dari 0.")
		return
	}
	//Cek Kondisi

	cekKondisi(5)

}
