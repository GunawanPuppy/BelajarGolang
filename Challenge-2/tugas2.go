package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ? fmt.Sprintf ini template literal dalam golang
// cara pakainya fmt.Sprintf(`%tipedata`, nama variabel urut dari tipe data nya)
// ? fmt.Println ini seperti console.log()
func main() {
	//!Soal 1
	var text1 string = "Bootcamp"
	var text2 string = "Digital"
	var text3 string = "Skill"
	var text4 string = "Sanbercode"
	var text5 string = "Golang"
	var sentence string = fmt.Sprintf(`%s %s %s %s %s`, text1, text2, text3, text4, text5)
	fmt.Println(sentence)

	//!Soal 2
	halo := "Halo Dunia"
	fmt.Println(strings.Replace(halo, "Dunia", "Golang", 1))

	//!Soal 3
	var kataPertama = "saya"
	kataKedua := "senang"
	kataKetiga := "belajar"
	kataKeempat := "golang"

	kataKedua = strings.Title(kataKedua)
	kataKetiga = kataKetiga[0:6] + strings.ToUpper(kataKetiga[6:7])
	kataKeempat = strings.ToUpper(kataKeempat)

	var kalimat = fmt.Sprintf(`%s %s %s %s`, kataPertama, kataKedua, kataKetiga, kataKeempat)
	fmt.Println(kalimat)

	//!Soal 4
	angkaPertama := "8"
	angkaKedua := "5"
	angkaKetiga := "6"
	angkaKeempat := "7"

	firstNumber, err := strconv.Atoi(angkaPertama)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	secondNumber, err := strconv.Atoi(angkaKedua)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	thirdNumber, err := strconv.Atoi(angkaKetiga)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fourthNumber, err := strconv.Atoi(angkaKeempat)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jumlah := firstNumber + secondNumber + thirdNumber + fourthNumber
	fmt.Println(jumlah)

	//!Soal 5
	var soal5 = "halo halo bandung"
	var num = 2021
	soal5 = strings.Replace(soal5, "halo", "Hi", 2)
	output := fmt.Sprintf(`"%s" - %d`, soal5, num)
	fmt.Println(output)
}
