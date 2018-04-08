package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
)

type input struct {
	damage      int
	robotProgam []string
}

func calculateDamage(program []string) int {
	strength := 1
	d := 0
	for _, instruction := range program {
		if instruction == "C" {
			strength = strength * 2
		} else if instruction == "S" {
			d += strength
		}
	}
	return d
}

func run(testInput input) (int, error) {
	moveCount := 0
	program := testInput.robotProgam
	for true {
		currentDamage := calculateDamage(program)
		if currentDamage <= testInput.damage {
			break
		}
		// greedy
		canSwap := false
		for i := 0; i < len(program); i++ {
			if i > 0 && program[i] == "S" && program[i-1] != "S" {
				t := program[i-1]
				program[i-1] = program[i]
				program[i] = t
				canSwap = true
				moveCount++
				break
			}
		}
		if !canSwap {
			return -1, nil
		}
	}
	return moveCount, nil
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	testCount, err := strconv.Atoi(reader.Text())
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= testCount && reader.Scan(); i++ {
		parts := strings.Split(strings.TrimSpace(reader.Text()), " ")
		d, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatal(err)
		}
		moveCount, err := run(input{damage: d, robotProgam: strings.Split(strings.TrimSpace(parts[1]), "")})
		if err != nil {
			log.Fatal(err)
		}
		moveCountStr := strconv.Itoa(moveCount)
		if moveCount < 0 {
			moveCountStr = "IMPOSSIBLE"
		}
		fmt.Printf("Case #%d: %s\n", i, moveCountStr)
	}
}
