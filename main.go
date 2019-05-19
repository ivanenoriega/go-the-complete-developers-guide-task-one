package main

import "fmt"

func main() {
	var t [11]int
	for i := range t {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
