package main

import (
	"fmt"
	"strconv"
)

func main() {
	//!Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	pp, err := strconv.Atoi(panjangPersegiPanjang)
	if err != nil {
		fmt.Println(err)
		return
	}
	lp, err := strconv.Atoi(lebarPersegiPanjang)
	if err != nil {
		fmt.Println(err)
		return
	}
	as, err := strconv.Atoi(alasSegitiga)
	if err != nil {
		fmt.Println(err)
		return
	}
	ts, err := strconv.Atoi(tinggiSegitiga)
	if err != nil {
		fmt.Println(err)
		return
	}

	var luasPersegiPanjang int = pp * lp
	var kelilingPersegiPanjang int = 2 * (pp + lp)
	var luasSegitiga int = as * ts / 2
	fmt.Println(luasPersegiPanjang, kelilingPersegiPanjang, luasSegitiga)

	//! Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiDoe < 50 {
		fmt.Println("Doe mendapatkan indeks E")
	} else if nilaiDoe < 60 {
		fmt.Println("Doe mendapatkan indeks D")
	} else if nilaiDoe < 70 {
		fmt.Println("Doe mendapatkan indeks C")
	} else if nilaiDoe < 80 {
		fmt.Println("Doe mendapatkan indeks B")
	} else {
		fmt.Println("Doe mendapatkan indeks A")
	}

	if nilaiJohn < 50 {
		fmt.Println("John mendapatkan indeks E")
	} else if nilaiJohn < 60 {
		fmt.Println("John mendapatkan indeks D")
	} else if nilaiJohn < 70 {
		fmt.Println("John mendapatkan indeks C")
	} else if nilaiJohn < 80 {
		fmt.Println("John mendapatkan indeks B")
	} else {
		fmt.Println("John mendapatkan indeks A")
	}

	//!Soal 3
	var tanggal = 17
	var bulan = 8
	var tahun = 1945
	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "July"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan"
	}
	fmt.Println(tanggal, namaBulan, tahun)

	//!Soal 4
	var year int = 1994
	if(year < 1965){
		fmt.Println("Baby Boomer")
	}else if(year < 1980){
		fmt.Println("Generasi X")
	}else if (year < 1995){
		fmt.Println("Generasi Y (Millenials)")
	}else if (year <2016){
		fmt.Println("Generasi Z")
	}
}
