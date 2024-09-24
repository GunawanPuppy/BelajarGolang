package main 
// import (
// 	"fmt"
// )

// func endApp(){
// 	fmt.Println("End App")
// 	// nangkep pesan error dengan recover
// 	message := recover()
// 	fmt.Println("Terjadi error", message)
// }

// func runApp(error bool){
// 	defer endApp()
// 	if error{
// 		panic("ERROR")
// 	}
// }

// func main(){
// 	runApp(false)
// }