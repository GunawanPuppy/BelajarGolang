package main 

// import (
// 	"fmt"
// 	// "time"
// )

// // func say(s string){
// // 	for i :=0; i< 5; i++{
// // 		time.Sleep(1000 * time.Millisecond)
// // 		fmt.Println(s)
// // 	}
// // }
// // func cetak(ch chan int){
// // 	fmt.Println("ini dari goroutine...")
// // 	//!mengirim nilai 10 ke ch
// // 	ch <- 10
// // }

// func print(ch chan <- int){
// 	for index :=0; index<10; index++{
// 		ch<- index
// 	}
// 	close(ch)
// }

// func main(){
// 	//dengan ini maka kadang bisa world bisa hello terlebih dahulu
// 	//karena function say world berjalan di thread lain
// 	// go say("world")
// 	// say("hello")

// 	//!membuat sebuah channel yg dapat menerima dan mengirim data tipe integer
// 	//!angka merupakan variabel yang menyimpan referensi channel 
// 	// angka:= make(chan int)
// 	// //! menjalankan fungsi cetak dari go routine
// 	// go cetak(angka)
// 	// //!menunggu nilai terisi ke channel angka lalu disimpan di variabel nilai
// 	// //!menerima nilai dari channel
// 	// nilai:= <-angka
// 	// fmt.Println("nilai channel integer:", nilai)
// 	// fmt.Println("ini dari function main...")

// 	//?
// 	ch := make(chan int)
// 	go print(ch)

// 	for {
// 		data, ok:= <- ch
// 		if ok == false{
// 			break
// 		}
// 	fmt.Printf("Data diterima %v\n", data)
// 	}
// }