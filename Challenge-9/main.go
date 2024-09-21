package main

import (
	"fmt"
	"Challenge-9/library"
)

func main() {
	segitiga := library.SegitigaSamaSisi{Alas: 4, Tinggi: 5}
	library.ShowBangunDatar(segitiga)

	persegi := library.PersegiPanjang{Panjang: 4, Lebar: 6}
	library.ShowBangunDatar(persegi)

	tabungObj := library.Tabung{JariJari: 7, Tinggi: 10}
	library.ShowBangunRuang(tabungObj)

	balokObj := library.Balok{Panjang: 4, Lebar: 3, Tinggi: 5}
	library.ShowBangunRuang(balokObj)

	//!Soal 2
	myPhone := library.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	var info library.PhoneInfo = myPhone
	fmt.Println(info.GetInfo())

	//!Soal 3
	fmt.Println(library.LuasPersegi(4, true))
	fmt.Println(library.LuasPersegi(8, false))
	fmt.Println(library.LuasPersegi(0, true))
	fmt.Println(library.LuasPersegi(0, false))

	//!Soal 4
	var prefix interface{} = "hasil penjumlahan dari"
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Type assertion
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)
	total := 0

	// Hitung total penjumlahan
	for _, value := range angkaPertama {
		total += value
	}
	for _, value2 := range angkaKedua {
		total += value2
	}

	// Format menghasilkan output yang diinginkan
	hasil := fmt.Sprintf("%s %d + %d + %d + %d = %d", prefix, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
	fmt.Println(hasil)
}
