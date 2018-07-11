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

const alive rune = 'X'
const dead rune = '.'

func main() {
	iterations := flag.Int("iterations", 20, "number of iterations")
	//rule := flag.Int64("rule", 110, "rule number")
	cells := flag.Int("cells", 32, "number of cells")
	flag.Parse()
	cellLine := make([]rune, *cells, *cells)
	initialize(cellLine)
	for i := 0; i < *iterations; i++ {
		printLine(cellLine)
		update(cellLine)
	}



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

func update(cells []rune) {

	previous := make([]rune, len(cells), cap(cells))
	copy(previous, cells)
	for i := range previous {
		var left rune
		var center rune
		var right rune
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

func applyRule(left, center, right rune) rune{
	a := string(alive)
	d := string(dead)
	state := alive
	pattern := string(string(left) + string(center) + string(right))
	switch pattern {
	case string(a+a+a), string(a+d+d), string(d+d+d):
		state = dead
	}
	return state
}

func printLine(cells []rune) {
	for _,value := range cells {
		fmt.Print(string(value))
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


