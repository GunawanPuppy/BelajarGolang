package main

import (
	"fmt"
	"sort"
	"sync"
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
}
