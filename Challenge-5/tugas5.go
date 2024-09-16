package main

import (
	"fmt"
	"strings"
)

// !Soal 1
func LuasPersegiPanjang(panjang int, lebar int) {
	fmt.Println(panjang * lebar)
}

func kelilingPersegiPanjang(panjang int, lebar int) {
	fmt.Println(2 * (panjang + lebar))
}
func volumeBalok(panjang int, lebar int, tinggi int) {
	fmt.Println(panjang * lebar * tinggi)
}

func main() {
	panjang := 12
	lebar := 4
	tinggi := 8
	LuasPersegiPanjang(panjang, lebar)
	kelilingPersegiPanjang(panjang, lebar)
	volumeBalok(panjang, lebar, tinggi)
	fmt.Println(introduce("Sarah", "perempuan", "penulis", "30"))
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)
}

// ! Soal 2
//predefined function
func introduce(name string, gender string, job string, age string) (kalimat string) {
	if gender == "laki-laki" {
		fmt.Printf(`Pak %s adalah seorang %s yang berusia %s tahun`, name, job, age)
	} else {
		fmt.Printf(`Bu %s adalah seorang %s yang berusia %s tahun`, name, job, age)
	}
	return kalimat
}

// !Soal 3
//Variadic Function
func buahFavorit(name string, buah ...string)(output string) {
	//separator pertama nambahin " + hasil string join + terakhir nambahin "
	buahList := "\"" + strings.Join(buah, "\", \"") + "\""
	

	output = fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, buahList)
	return output
}

// !Soal 4
//Closure Function



