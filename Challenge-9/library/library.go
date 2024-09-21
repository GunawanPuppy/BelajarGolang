package library

import (
	"fmt"
	"math"
	"strings"
)

// !Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64((b.Panjang*b.Lebar)+(b.Panjang*b.Tinggi)+(b.Lebar*b.Tinggi))
}

func ShowBangunDatar(bd HitungBangunDatar) {
	fmt.Printf("Luas: %d\n", bd.Luas())
	fmt.Printf("Keliling: %d\n\n", bd.Keliling())
}

func ShowBangunRuang(br HitungBangunRuang) {
	//.2 menunjukkan jumlah digit setelah desimal ada 2
	fmt.Printf("Volume: %.2f\n", br.Volume())
	fmt.Printf("LuasPermukaan: %.2f\n\n", br.LuasPermukaan())
}

// !Soal 2
type Phone struct {
	Name, Brand string
	Year      int
	Colors      []string
}

type PhoneInfo interface {
	GetInfo() string
}

func (p Phone) GetInfo() string {
	colors := strings.Join(p.Colors, ", ")
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %s", p.Name, p.Brand, p.Year, colors)
}
// !Soal 3
func LuasPersegi(Sisi int, InfoDetail bool) interface{} {
	//jika Sisi 0  dan InfoDetail true
	if Sisi == 0 && InfoDetail {
		return "Maaf anda belum menginput Sisi dari persegi"
	}
	//jika Sisi 0 dan InfoDetail false
	if Sisi == 0 && !InfoDetail {
		return nil
	}

	luas := Sisi * Sisi
	//jika info detail true
	if InfoDetail {
		return fmt.Sprintf("Luas Persegi dengan Sisi %d cm adalah %d cm", Sisi, luas)
	}
	// jika info detail false
	return luas
}