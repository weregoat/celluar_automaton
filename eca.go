package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
	"strconv"
	"os"
	"strings"
	"log"
)

const alive = 'X'
const dead = ' '

type transitionRule struct {
	configuration string
	state rune
}

/* Rules */
var rules [8]transitionRule

/* Wolfram Code */
var ruleNumber int



func main() {
	iterations := flag.Int("iterations", 20, "number of iterations")
	flag.IntVar(&ruleNumber,"rule", 110, "rule number")
	cells := flag.Int("cells", 32, "number of cells")
	flag.Parse()
	if ruleNumber < 0 || ruleNumber > 255 {
		log.Fatal(fmt.Sprintf("Invalid rule number %d", ruleNumber))
		os.Exit(1)
	}
	populateTable(ruleNumber)
	cellLine := make([]rune, *cells, *cells)
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
		pattern := fmt.Sprintf("%s%s%s", string(left),string(center),string(right))
		rule := rules[toInt(pattern)]
		cells[i] = rule.state
	}
}


func printLine(cells []rune) {
	for _,value := range cells {
		fmt.Print(string(value))
	}
	fmt.Print("\n")
}


func populateTable(ruleNumber int) {

	/* Generate the states as a binary string */
	binaryString := toString(ruleNumber, 8)
	for i := range rules {
		pattern := toString(i, 3)
		state := binaryString[(len(binaryString)-i)-1]
		rule := transitionRule{pattern, rune(state)}
		rules[i] = rule
	}

}

func toString(integer int, len int) string {
	binary := fmt.Sprintf("%0[1]*s",len,strconv.FormatInt(int64(integer),2))
	firstRound := strings.Replace(binary, "1", string(alive),-1)
	secondRound := strings.Replace(firstRound, "0", string(dead), -1)
	return secondRound
}

func toInt(pattern string) int {
	firstRound := strings.Replace(pattern, string(alive), "1",-1)
	binary := strings.Replace(firstRound, string(dead), "0", -1)

	value,err := strconv.ParseInt(binary,2,32)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not convert pattern %s to int", binary))
		os.Exit(99)
	}
	return int(value)

}


