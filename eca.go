package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
)

type cell struct {
	left int
	center int
	right int
}

const alive = 'X'
const dead = ' '

func main() {
	iterations := flag.Int("iterations", 20, "number of iterations")
	//rule := flag.Int64("rule", 110, "rule number")
	cells := flag.Int("cells", 32, "number of cells")
	flag.Parse()
	cellLine := make([]int, *cells, *cells)
	i := 0
	initialize(cellLine)
	printLine(cellLine)
	i++
	for i < *iterations {
		update(cellLine)
		printLine(cellLine)
		i++
	}



}

func initialize(cells []int) {
	seed := rand.NewSource(time.Now().UnixNano())
	randomiser := rand.New(seed)
	for i := range cells {
		state := randomiser.Intn(2)
		if state == 0 {
			cells[i] = 0
		} else {
			cells[i] = 1
		}
	}
}

func update(cells []int) {

	previous := make([]int, len(cells), cap(cells))
	copy(previous, cells)
	for i := range previous {
		var left int
		var center int
		var right int
		/* Special cases */
		if i == 0 {
			left = previous[len(previous)-1]
		} else {
			left = previous[i-1]
		}
		if i == len(previous)-1 {
			right = previous[0]
		} else {
			right = previous[i+1]
		}
		center = previous[i]
		cells[i] = applyRule(left, center, right)
	}
}

func applyRule(left, center, right int) int{
	state := 1
	pattern := fmt.Sprintf("%d%d%d", left,center,right)
	switch pattern {
	case "111", "100", "000":
		state = 0
	}
	return state
}

func printLine(cells []int) {
	for _,value := range cells {
		symbol := dead
		if value == 1 {
			symbol = alive
		}
		fmt.Print(string(symbol))
	}
	fmt.Print("\n")
}

/*
func generateTable(rule int64) map[string]rune{
	/* Initialize the transformation table */
	/* The key is the LCR cell combination */
	/* The value is the center next value */
	/*
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
*/


