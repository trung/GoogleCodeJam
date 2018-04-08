package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func troubleSort(a []int) []int {
	for true {
		noMore := true
		for i := 0; i < len(a) - 2; i++ {
			//fmt.Println("inspect", a[i], a[i+1], a[i+2])
			if a[i] > a[i+2] {
				t := a[i]
				a[i] = a[i+2]
				a[i+2] = t
				noMore = false
			}
		}
		//fmt.Println(a)
		if noMore {
			break
		}
	}
	return a
}

func verify(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testCount, _ := strconv.Atoi(scanner.Text())
	for i := 1; i <= testCount && scanner.Scan(); i++ {
		l, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		elements := strings.Split(scanner.Text(), " ")
		a := make([]int, 0)
		for j := 0; j < l; j++ {
			v, _ := strconv.Atoi(strings.TrimSpace(elements[j]))
			a = append(a, v)
		}
		a = troubleSort(a)
		errorIndex := verify(a)
		y := strconv.Itoa(errorIndex)
		if errorIndex < 0 {
			y = "OK"
		}
		fmt.Printf("Case #%d: %s\n", i, y)
	}
}
