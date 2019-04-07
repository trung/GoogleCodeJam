package main

import "fmt"

func main() {
	n := 50000
	for i := 0; i < 2*n - 2; i++ {
		if i % 2 == 0 {
			fmt.Print("E")
		} else {
			fmt.Print("S")
		}
	}
}
