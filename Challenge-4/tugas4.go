package main

import (
	"fmt"
)

func main() {
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
}
