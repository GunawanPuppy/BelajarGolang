package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func showPhone(i int, value string, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(1 * time.Second)
	output:= fmt.Sprintf("%d. %s", i+1, value)
	fmt.Println(output)
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
}
