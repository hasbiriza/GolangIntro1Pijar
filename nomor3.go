package main

import "fmt"

func condition(rata_rata int) string {
	if rata_rata >= 90 && rata_rata <= 100 {
		return "A"
	} else if rata_rata >= 80 && rata_rata < 90 {
		return "B"
	} else if rata_rata >= 70 && rata_rata < 80 {
		return "C"
	} else if rata_rata >= 60 && rata_rata < 70 {
		return "D"
	} else {
		return "E"
	}
}

func validasiDanPerhitungan(mtk, bindo, bing, ipa int) interface{} {
	if mtk < 0 || bindo < 0 || bing < 0 || ipa < 0 {
		return "Semua nilai harus diisi. Silakan coba lagi."
	}

	hasil := (mtk + bindo + bing + ipa) / 4
	return hasil
}

func main() {
	hasilvalidasi_perhitungan := validasiDanPerhitungan(90, 80, 41, 63)

	// Melakukan tipe assertion
	if Sumperhitungan, ok := hasilvalidasi_perhitungan.(int); ok {
		fmt.Println("Rata-Rata =", Sumperhitungan)
		grade := condition(Sumperhitungan)
		fmt.Println("Grade =", grade)
	} else {
		// Handle jika tipe assertion gagal
		fmt.Println("Error: Hasil tidak valid, Cek Nilai Yang dimasukkan !.")
	}

}

// func validasi(mtk, bindo, bing, ipa int) string {
// 	if mtk < 0 || bindo < 0 || bing < 0 || ipa < 0 {
// 		return "Semua nilai harus diisi. Silakan coba lagi."
// }

// func perhitungan(mtk, bindo, bing, ipa int) (hasil int) {
// 	hasil = mtk + bindo + bing + ipa/4
// 	return
// }

// // Menentukan grade pakai Switch
// var grade string
// switch {
// case rataRata >= 90 && rataRata <= 100:
// 	grade = "A"
// case rataRata >= 80:
// 	grade = "B"
// case rataRata >= 70:
// 	grade = "C"
// case rataRata >= 60:
// 	grade = "D"
// default:
// 	grade = "E"
// }
