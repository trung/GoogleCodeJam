package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	addition = [10][2]byte{
		{0, 0}, // 0
		{0, 1}, // 1
		{1, 1}, // 2
		{1, 2}, // 3
		{1, 3}, // 4
		{2, 3}, // 5
		{1, 5}, // 6
		{1, 6}, // 7
		{2, 6}, // 8
		{1, 8}, // 9
	}
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	testCount, err := strconv.Atoi(reader.Text())
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= testCount && reader.Scan(); i++ {
		input := reader.Text()
		number1 := make([]byte, len(input))
		number2 := make([]byte, len(input))
		for idx, c := range input {
			digit := c - '0'
			s := addition[digit]
			number1[idx] = s[0]
			number2[idx] = s[1]
		}
		fmt.Printf("Case #%d: %s %s\n", i, toNumberStr(number1), toNumberStr(number2))
	}
}

func toNumberStr(digits []byte) string {
	hasLeadingZeros := true
	newDigits := make([]byte, 0, len(digits))
	for _, i := range digits {
		if hasLeadingZeros && i == 0 {
			continue
		} else {
			hasLeadingZeros = false
		}
		newDigits = append(newDigits, i+'0')
	}
	return string(newDigits)
}
