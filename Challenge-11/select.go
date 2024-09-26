package main

// import (
// 	"fmt"
// )

// func getAverage(numbers []int, ch chan float64) {
// 	var sum = 0
// 	for _, value := range numbers {
// 		sum += value

// 	}
// 	//mengirim data ke channel
// 	ch <- float64(sum) / float64(len(numbers))
// }

// func getMax(numbers []int, ch chan int) {
// 	var max = numbers[0]
// 	for _, value := range numbers {
// 		if max < value {
// 			max = value

// 		}
// 	}
// 	//mengirim data ke channel
// 	ch <- max
// }

// func main() {
// 	numbers := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
// 	fmt.Println("numbers", numbers)
// 	//membuat channel 1
// 	var ch1 = make(chan float64)
// 	//mengeksekusi go routine dengan mengirimkan numbers dan ch1 sebagai parameter di function getAverage
// 	go getAverage(numbers, ch1)
// 	//membuat channel2
// 	var ch2 = make(chan int)
// 	//mengeksekusi go routine dengan mengirimkan numbers dan ch 2 sebagai parameter di function getMax
// 	go getMax(numbers, ch2)
// 	// \t memberi tab
// 	for i := 0; i < 2; i++ {
// 		select {
// 		// case avg  menerima nilai ch1 ke variabel avg
// 		case avg := <-ch1:
// 			fmt.Printf("Avg \t: %.2f \n", avg)
// 			// case max menerima nilai ch2 ke variabel max
// 		case max := <-ch2:
// 			fmt.Printf("Max \t: %d \n", max)
// 		}
// 	}
// }
