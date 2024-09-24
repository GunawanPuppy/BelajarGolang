package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// !Soal 1
func jalan(kalimat string, tahun int) {
	output := fmt.Sprintf("%s %d", kalimat, tahun)
	fmt.Println(output)
}

func runFunction(kalimat string, tahun int) {
	defer jalan(kalimat, tahun)
}

// !Soal 2
func kelilingSegitigaSamaSisi(sisi int, infoDetail bool) interface{} {
	keliling := sisi * 3
	if sisi == 0 && !infoDetail {
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	if sisi == 0 && infoDetail {
		return "Maaf anda belum menginput sisi dari segitiga sama sisi"
	}
	if infoDetail {
		output := fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling)
		return output
	}
	return keliling
}

func message() {
	pesan := recover()
	if pesan != nil {
		fmt.Println(pesan, "ini didapat dari panic")
	}
}

// !Soal 3
func tambahAngka(nilai int, total *int) {
	*total += nilai
}

// Fungsi untuk mencetak angka
func cetakAngka(total *int) {
	fmt.Println("Total angka:", *total)
}

//!Soal 4
func addPhone(phones *[]string) {
	*phones = append(*phones, "Xiaomi", "Iphone")
}

func listPhone(phones *[]string) {
	sort.Strings(*phones)
	for key, value := range *phones {
		output := fmt.Sprintf(`%d. %s`, key+1, value)
		fmt.Println(output)
		//seperti setTimeout pada javascript
		time.Sleep(1 * time.Second)
	}
}

//!Soal 5
func luasLingkaran(radius float64){
	//Pow = pangkat
	hasil := math.Pi * math.Pow(radius,2)
	fmt.Println(hasil)
}
func kelilingLingkaran(radius float64){
	hasil := math.Pi * 2 * radius
	fmt.Println(hasil)
}

func main() {
	angka := 1

	defer cetakAngka(&angka)
	defer message()

	//!Soal 1
	sentence := "Golang Backend Development"
	year := 2021
	runFunction(sentence, year)

	//!Soal 3
	// Menyimpan hasil penambahan ke variabel lokal

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
	//!Soal 4
	var phones = []string{}
	addPhone(&phones)
	listPhone(&phones)

	//!Soal 5
	var jariJari float64 = 7
	luasLingkaran(jariJari)
	kelilingLingkaran(jariJari)
	jariJari = 10
	luasLingkaran(jariJari)
	kelilingLingkaran(jariJari)
	jariJari = 15
	luasLingkaran(jariJari)
	kelilingLingkaran(jariJari)

	//!Soal 6
	
	panjang := flag.Float64("panjang",0,"Masukkan nilai panjang persegi panjang")
	lebar := flag.Float64("lebar",0,"Masukkan nilai lebar persergi panjang")
	flag.Parse()

	luas := (*panjang) * (*lebar)
	keliling := 2 * ((*panjang) + (*lebar))

	fmt.Printf("Luas persegi panjang: %.2f\n", luas)
	fmt.Printf("Keliling persegi panjang : %.2f\n", keliling)

	//!Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

}
