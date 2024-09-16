package main

import (
	"fmt"
)

// !Soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func printBuah(buahs []buah) {
	for _, value := range buahs {
		biji := "Ada"
		if !value.adaBijinya {
			biji = "Tidak"
		}
		fmt.Printf("Nama:%s\nWarna: %s\nAdaBijinya: %s\nHarga %d\n\n", value.nama, value.warna, biji, value.harga)
	}
}

// !Soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (a segitiga) luasSegitiga() {
	luasTiga := a.alas * a.tinggi / 2
	fmt.Println(luasTiga)
}
func (b persegi) luasPersegi() {
	luasPersegis := b.sisi * b.sisi
	fmt.Println(luasPersegis)
}
func (c persegiPanjang) luasPersegiPanjang() {
	luasPersegiPanjangs := c.panjang * c.lebar
	fmt.Println(luasPersegiPanjangs)
}

// !Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// menggunakan pointer supaya setelah ditambahkan objek asli juga terganti valuenya
func (p *phone) addColor(newColor string) {
	p.colors = append(p.colors, newColor)
}

// !Soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, genre string, duration int, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, newMovie)
}

func showDataFilm(dataFilm []movie) {
	for i, value := range dataFilm {
		fmt.Printf("%d. title:%s\n duration:%d menit\n genre:%s\n year:%d\n\n", i+1, value.title, value.duration, value.genre, value.year)
	}
}

func main() {
	//!Soal 1
	buahs := []buah{
		{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000},
		{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000},
		{nama: "Semangka", warna: "Hijau&Merah", adaBijinya: true, harga: 10000},
		{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 5000},
	}
	printBuah(buahs)
	//!Soal 2
	var alas = segitiga{4, 6}
	alas.luasSegitiga()
	var sisi = persegi{4}
	sisi.luasPersegi()
	var pl = persegiPanjang{3, 4}
	pl.luasPersegiPanjang()
	//!Soal 3
	myPhone := phone{
		name:  "Iphone X",
		brand: "Apple",
		year:  2019,
		colors: []string{
			"Gold", "Red",
		},
	}
	myPhone.addColor("Blue")
	fmt.Println(myPhone)
	//!Soal 4
	var dataFilm = []movie{}
	tambahDataFilm("LOTR", "action", 120, 1999, &dataFilm)
	tambahDataFilm("avenger", "action", 120, 2019, &dataFilm)
	tambahDataFilm("spiderman", "action", 120, 2004, &dataFilm)
	tambahDataFilm("juon", "horror", 120, 2004, &dataFilm)
	showDataFilm(dataFilm)
}
