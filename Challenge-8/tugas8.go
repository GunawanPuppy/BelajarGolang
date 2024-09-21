package main

import (
	"fmt"
	"math"
	"strings"
)

// !Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.panjang*b.tinggi)+(b.lebar*b.tinggi))
}

func showBangunDatar(bd hitungBangunDatar) {
	fmt.Printf("Luas: %d\n", bd.luas())
	fmt.Printf("Keliling: %d\n\n", bd.keliling())
}

func showBangunRuang(br hitungBangunRuang) {
	//.2 menunjukkan jumlah digit setelah desimal ada 2
	fmt.Printf("Volume: %.2f\n", br.volume())
	fmt.Printf("LuasPermukaan: %.2f\n\n", br.luasPermukaan())
}

// !Soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneInfo interface {
	getInfo() string
}

func (p phone) getInfo() string {
	colors := strings.Join(p.colors, ", ")
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %s", p.name, p.brand, p.year, colors)
}

// !Soal 3
func luasPersegi(sisi int, infoDetail bool) interface{} {
	//jika sisi 0  dan infoDetail true
	if sisi == 0 && infoDetail {
		return "Maaf anda belum menginput sisi dari persegi"
	}
	//jika sisi 0 dan infoDetail false
	if sisi == 0 && !infoDetail {
		return nil
	}

	luas := sisi * sisi
	//jika info detail true
	if infoDetail {
		return fmt.Sprintf("Luas Persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	// jika info detail false
	return luas
}

func main() {
	segitiga := segitigaSamaSisi{alas: 4, tinggi: 5}
	showBangunDatar(segitiga)

	persegi := persegiPanjang{panjang: 4, lebar: 6}
	showBangunDatar(persegi)

	tabungObj := tabung{jariJari: 7, tinggi: 10}
	showBangunRuang(tabungObj)

	balokObj := balok{panjang: 4, lebar: 3, tinggi: 5}
	showBangunRuang(balokObj)
	//!Soal 2
	myPhone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	var info phoneInfo = myPhone
	fmt.Println(info.getInfo())

	//!Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//!Soal 4
	var prefix interface{} = "hasil penjumlahan dari"
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	//Type assertion 
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)
	total := 0

	//Hitung total penjumlahan
	for _, value  := range angkaPertama {
		total += value
	}
	for _, value2 := range angkaKedua {
		total += value2
	}

	//Format menghasilkan output yang diinginkan
	hasil := fmt.Sprintf("%s %d + %d + %d + %d = %d",prefix,angkaPertama[0],angkaPertama[1],angkaKedua[0],angkaKedua[1],total)
	fmt.Println(hasil)
}
