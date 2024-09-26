package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
	// "time"
)



func showPhone(i int, value string, wg *sync.WaitGroup){
	defer wg.Done()
	// time.Sleep(1 * time.Second)
	output:= fmt.Sprintf("%d. %s", i+1, value)
	fmt.Println(output)
}

func getMovies(moviesChannel chan string,movies ...string){
	for i, value := range movies {
		moviesChannel <- fmt.Sprintf("%d. %s", i+1,value)
	}
	close(moviesChannel)
}
//!Soal 3
func luasLingkaran(radius float64, ch chan<- float64){
	luas := math.Pi * radius *radius
	ch <- luas
}
func kelilingLingkaran(radius float64, ch chan<- float64){
	keliling := math.Pi *2 *radius
	ch <- keliling
}
func volumeTabung(radius float64, tinggi float64,ch chan<- float64){
	volume := math.Pi *radius *radius *tinggi
	ch <- volume
}
//!Soal 4 
func luasPP(panjang int, lebar int, ch chan int){
	luas := panjang * lebar
	ch <- luas
}

func kelilingPP(panjang int, lebar int, ch chan int){
	keliling := 2 * (panjang+lebar)
	ch <- keliling
}

func volumeBalok(panjang int, lebar int, tinggi int, ch chan int){
	volume := panjang * lebar * tinggi
	ch <- volume
}

func main() {
	//!Soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)
	
	var wg sync.WaitGroup

	for i, value := range phones {
		wg.Add(1)
		go showPhone(i,value,&wg)
		wg.Wait()
	}
	//!Soal 2
	var movies = []string{"Harry Potter","LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)

	go getMovies(moviesChannel,movies...)
	fmt.Println("List Movies:")
	for value := range moviesChannel{
		fmt.Println(value)
	}
	//!Soal 3 
	radius:= []float64{8,14,20}
	var tinggi float64 = 10
	ch := make(chan float64)
	for _, value := range radius {
		go luasLingkaran(value,ch)
		go kelilingLingkaran(value,ch)
		go volumeTabung(value,tinggi,ch)
	}

	for i := 0; i < len(radius)*3; i++ {
		result := <- ch
		time.Sleep(1*time.Second)
		fmt.Println(result, "ini result")
	}
	//!Soal 4
	length := 10
	width := 10
	height := 3
	area := make(chan int)
	go luasPP(length,width,area)
	perimeter := make(chan int)
	go kelilingPP(length,width,perimeter)
	volume := make(chan int)
	go volumeBalok(length,width,height,volume)
	for i := 0; i < 3; i++ {
		select{
		case areaPP := <- area:
			fmt.Printf("Luas persegi panjang : %d\n", areaPP)
		case perimeterPP := <- perimeter:
			fmt.Printf("Keliling persegi panjang: %d\n", perimeterPP)
		case blockVolume := <- volume:
			fmt.Printf("Volume Balok : %d\n", blockVolume)
		}
	}
}
