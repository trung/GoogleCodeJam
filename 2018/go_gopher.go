package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	a     int
	d     int
	field [][]int
	nextX []int
	nextY []int
	step  int
)

func prepare(x, y int) {
	fmt.Printf("%d %d\n", x, y)
}

func filled(x, y int) bool {
	for i := x; i < x+2; i++ {
		for j := y; j < y+2; j++ {
			if field[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func tooManyAttempts(x, y int) bool {
	attempts := 0
	for i := x; i < x+2; i++ {
		for j := y; j < y+2; j++ {
			attempts += field[i][j]
		}
	}
	return attempts > 10
}

func nextTarget(lastTargetX, lastTargetY int) (int, int) {
	if filled(lastTargetY-1, lastTargetY-1) || tooManyAttempts(lastTargetX - 1, lastTargetY - 1) {
		step++
		if step >= len(nextX) {
			step = 0
		}
		return nextX[step], nextY[step]
	} else {
		return lastTargetX, lastTargetY
	}
}

func resetField() {
	field = make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		field[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			field[i][j] = 0
		}
	}
}

func determinSize() (int, int) {
	for i := 3; i < 40; i++ {
		if i*i > a {
			d = i
			break
		}
	}
	// fmt.Fprintf(os.Stderr, "Size: %d\n", d)
	nextX = make([]int, 0)
	nextY = make([]int, 0)
	step = 0
	x := 2
	for x <= d-1 {
		y := 2
		for y <= d-1 {
			nextX = append(nextX, x)
			nextY = append(nextY, y)
			y += 2
			if y == d {
				y--
			}
		}
		x += 2
		if x >= d {
			x--
		}
	}
	// fmt.Fprintln(os.Stderr, nextX, nextY)
	return nextX[0], nextY[0]
}

func printField() {
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			fmt.Fprintf(os.Stderr, "%-4d", field[i][j])
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testCount, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < testCount; i++ {
		resetField()
		scanner.Scan()
		a, _ = strconv.Atoi(scanner.Text())
		targetX, targetY := determinSize()
		success := false
		error := false
		for true {
			prepare(targetX, targetY)
			scanner.Scan()
			xy := strings.Split(scanner.Text(), " ")
			preparedX, _ := strconv.Atoi(xy[0])
			preparedY, _ := strconv.Atoi(xy[1])
			success = preparedX == 0 && preparedY == 0
			error = preparedX == -1 && preparedY == -1
			if success || error {
				break
			}
			field[preparedX-1][preparedY-1] += 1
			targetX, targetY = nextTarget(targetX, targetY)
		}
		if error {
			printField()
			break
		}
	}
}
