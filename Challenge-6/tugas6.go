package main

import (
	"fmt"
	"math"
)

var sentence string

func LuasLingkaran(jariJari *float64) float64 {
	// *melakukan dereferencing radius
	LuasLingkaran := math.Pi * *jariJari * *jariJari
	// fmt.Println(jariJari, "ini alamat memori")
	// fmt.Println(*jariJari,"ini ngambil value dari alamat memori")
	return LuasLingkaran
}

func KelilingLingkaran(radius *float64) float64 {
	KelilingLingkaran := math.Pi * 2 * *radius
	return KelilingLingkaran
}

func introduce(sentence *string, name *string, gender *string, profession *string, age *string) {
	if *gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", *name, *profession, *age)
	} else if *gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", *name, *profession, *age)
	}
}

func addBuah(buah *[]string) {
	*buah = append(*buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}
func listBuah(buah *[]string) {
	for key, value := range *buah {
		output := fmt.Sprintf(`%d. %s`, key, value)
		fmt.Println(output)
	}
}

func tambahDataFilm(dataFilm *[]map){

}

func main() {
	//!Soal 1
	radius := 7.0
	//mengambil value by reference radius
	fmt.Println(LuasLingkaran(&radius))
	fmt.Println(KelilingLingkaran(&radius))
	//!Soal 2

	name := "John"
	gender := "laki-laki"
	profession := "penulis"
	age := "30"
	introduce(&sentence, &name, &gender, &profession, &age)
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	//! Soal 3
	var buah = []string{}
	addBuah(&buah)
	// fmt.Println(buah)
	listBuah(&buah)
	
	//!Soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
}
