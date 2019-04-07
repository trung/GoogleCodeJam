package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	testCount, err := strconv.Atoi(reader.Text())
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= testCount; i++ {
		reader.Scan()
		size, _ := strconv.Atoi(reader.Text())
		reader.Scan()
		steps := reader.Text()
		row, col := uint16(0), uint16(0)
		herCords := make([][2]uint16, 2*size-1)
		herCords[0] = [2]uint16{0, 0}
		for idx, s := range steps {
			switch s {
			case 'E':
				col++
			case 'S':
				row++
			}
			herCords[idx+1] = [2]uint16{row, col}
		}
		myCords := make([][2]uint16, 2*size-1)
		myCords[0] = [2]uint16{0, 0}
		backtrack(0, uint16(size), herCords[:], myCords[:])
		mySteps := make([]byte, 2*size-2)
		for j, c := range myCords {
			if j == 0 {
				continue
			}
			s := byte('S')
			if c[0] == myCords[j-1][0] {
				s = 'E'
			}
			mySteps[j-1] = s
		}
		fmt.Printf("Case #%d: %s\n", i, string(mySteps))
	}
}

func backtrack(lastStep int, size uint16, herCords [][2]uint16, myCords [][2]uint16) bool {
	// fmt.Println(herCords[lastStep], "--", myCords[lastStep])
	if lastStep == int(2*size-2) {
		return true
	}
	found := false
	visitedSouth := false
	visitedEast := false
	myNextRowSouth, myNextColSouth, okSouth := canMove(lastStep, size, herCords[:], myCords[:], 1, 0)
	myNextRowEast, myNextColEast, okEast := canMove(lastStep, size, herCords[:], myCords[:], 0, 1)
	if okSouth && okEast {
		herNextRow, herNextCol := herCords[lastStep+1][0], herCords[lastStep+1][1]
		// select the cord nearer to hers
		distSouth := math.Abs(float64(int(myNextRowSouth)-int(herNextRow))) + math.Abs(float64(int(myNextColSouth)-int(herNextCol)))
		distEast := math.Abs(float64(int(myNextRowEast)-int(herNextRow))) + math.Abs(float64(int(myNextColEast)-int(herNextCol)))
		// fmt.Println(distSouth, "--", distEast)
		if distEast > distSouth {
			visitedSouth = true
			myCords[lastStep+1] = [2]uint16{myNextRowSouth, myNextColSouth}
			found = backtrack(lastStep+1, size, herCords[:], myCords[:])
		}
		if !found {
			visitedEast = true
			myCords[lastStep+1] = [2]uint16{myNextRowEast, myNextColEast}
			found = backtrack(lastStep+1, size, herCords[:], myCords[:])
		}
	}
	if !found {
		if !visitedSouth && okSouth {
			myCords[lastStep+1] = [2]uint16{myNextRowSouth, myNextColSouth}
			found = backtrack(lastStep+1, size, herCords[:], myCords[:])
		}
		if !visitedEast && okEast {
			myCords[lastStep+1] = [2]uint16{myNextRowEast, myNextColEast}
			found = backtrack(lastStep+1, size, herCords[:], myCords[:])
		}
	}
	return found
}

func canMove(lastStep int, size uint16, herCords [][2]uint16, myCords [][2]uint16, deltaRow uint16, deltaCol uint16) (uint16, uint16, bool) {
	herLastRow, herLastCol, herNextRow, herNextCol := herCords[lastStep][0], herCords[lastStep][1], herCords[lastStep+1][0], herCords[lastStep+1][1]
	myLastRow, myLastCol, myNewRow, myNewCol := myCords[lastStep][0], myCords[lastStep][1], myCords[lastStep][0]+deltaRow, myCords[lastStep][1]+deltaCol
	same := myLastRow == herLastRow && myLastCol == herLastCol && myNewRow == herNextRow && myNewCol == herNextCol
	outOfBound := myNewRow == size || myNewCol == size
	return myNewRow, myNewCol, !same && !outOfBound
}
