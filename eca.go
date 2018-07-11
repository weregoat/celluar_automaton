package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
	"strconv"
)

type cell struct {
	left int
	center int
	right int
}

const alive rune = '1'
const dead rune = '0'

func main() {
	iterations := flag.Int("iterations", 20, "number of iterations")
	rule := flag.Int64("rule", 110, "rule number")
	cells := flag.Int("cells", 10, "number of cells")
	flag.Parse()
	cellLine := make([]rune, *cells, *cells)
	initialize(cellLine)
	fmt.Println(cellLine)
	for i := 0; i < *iterations; i++ {
		newLine := update(cellLine)
		fmt.Println(newLine)
		cellLine = newLine
	}
	generateTable(*rule)


}

func initialize(cells []rune) {
	seed := rand.NewSource(time.Now().UnixNano())
	randomiser := rand.New(seed)
	for i := range cells {
		state := randomiser.Intn(2)
		if state == 0 {
			cells[i] = dead
		} else {
			cells[i] = alive
		}
	}
}

func update(cells []rune) []rune{

	next := make([]rune, len(cells), cap(cells))
	copy(next, cells)
	for i := range cells {
		var left rune
		var center rune
		var right rune
		/* Special cases */
		if i == 0 {
			left = cells[len(cells)-1]
		} else {
			left = cells[i-1]
		}
		if i == len(cells)-1 {
			right = cells[0]
		} else {
			right = cells[i+1]
		}
		center = cells[i]
		next[i] = applyRule(left, center, right)
	}
	return next
}

func applyRule(left, center, right rune) rune{
//	newState := (center+right+center*right+left*center*right)%2
//	return newState
	return dead
}

func generateTable(rule int64) map[string]rune{
	/* Initialize the transformation table */
	/* The key is the LCR cell combination */
	/* The value is the center next value */
	table := map[string]rune {
		"111": dead,
		"110": dead,
		"101": dead,
		"100": dead,
		"011": dead,
		"010": dead,
		"001": dead,
		"000": dead,
	}
	binaryString := strconv.FormatInt(rule,2)
	for i := len(binaryString); i >= 0; i-- {
		value := binaryString[i]
		
	}
	fmt.Println(string(binaryString[0]))
	return table

}



