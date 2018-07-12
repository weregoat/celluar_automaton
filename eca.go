package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
	"strconv"
	"os"
)

type cell struct {
	left int
	center int
	right int
}

const alive = 'X'
const dead = ' '


/* Patterns in sequence */
var patterns = []string{"111","110","101","100","011","010","001","000"}
var rule int64

/* Empty table with pattern as key and the value as state */
var table = map[string]uint8 {
	"111":0,
	"110":0,
	"101":0,
	"100":0,
	"011":0,
	"010":0,
	"001":0,
	"000":0,
}

func main() {
	iterations := flag.Int("iterations", 20, "number of iterations")
	flag.Int64Var(&rule,"rule", 110, "rule number")
	cells := flag.Int("cells", 32, "number of cells")
	flag.Parse()
	if rule < 0 || rule > 255 {
		fmt.Println("Invalid rule number")
		os.Exit(1)
	}
	table = populateTable(rule)
	cellLine := make([]uint8, *cells, *cells)
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

func initialize(cells []uint8) {
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

func update(cells []uint8) {

	previous := make([]uint8, len(cells), cap(cells))
	copy(previous, cells)
	for i := range previous {
		var left uint8
		var center uint8
		var right uint8
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
		pattern := fmt.Sprintf("%d%d%d", left,center,right)
		cells[i] = table[pattern]
	}
}


func printLine(cells []uint8) {
	for _,value := range cells {
		symbol := dead
		if value == 1 {
			symbol = alive
		}
		fmt.Print(string(symbol))
	}
	fmt.Print("\n")
}


func populateTable(rule int64) map[string]uint8 {

	/* Generate the rule as a binary string */
	binaryString := fmt.Sprintf("%08s",strconv.FormatInt(rule,2))
	for i := 0; i < len(binaryString); i++ {
		value := rune(binaryString[i])
		if value == '1' {
			table[patterns[i]] = 1
		}
	}
	return table
}



