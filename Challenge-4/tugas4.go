package main

import (
	"fmt"
)

func main() {
	//! Soal 1
	num := 20
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%2 == 1 {
			fmt.Println(i, `- I Love Coding`)
		} else if i%2 == 0 {
			fmt.Println(i, `- Berkualitas`)
		} else if i%2 == 1 {
			fmt.Println(i, `- Santai`)
		}
	}
	//!Soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			//print # di tiap baris
			fmt.Print("#")
		}
		//mencetak baris baru setelah print pagar
		fmt.Println()
	}
	//!Soal 3
	var kalimat = [...]string{"aku","dan","saya","sangat","senang","belajar","golang"}
	
}
