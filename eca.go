package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
	"strconv"
	"os"
	"log"
)

const aliveStatus = '1'
const deadStatus = '0'


/* Rules */
var transitionRules map [string]rune

/* Wolfram Code */
var ruleNumber int



func main() {
	iterations := flag.Int("iterations", 20, "number of iterations")
	flag.IntVar(&ruleNumber,"rule", 110, "rule number")
	cells := flag.Int("cells", 32, "number of cells")
	deadSymbol := flag.String("dead", string(deadStatus), "Symbol representing dead cells")
	aliveSymbol := flag.String("alive", string(aliveStatus), "Symbol representing living cells")
	flag.Parse()
	if ruleNumber < 0 || ruleNumber > 255 {
		log.Fatal(fmt.Sprintf("Invalid rule number %d", ruleNumber))
		os.Exit(1)
	}
	transitionRules = createTransitionRules(ruleNumber)
	cellLine := make([]rune, *cells, *cells)
	i := 0
	initialize(cellLine)
	printLine(cellLine, *aliveSymbol, *deadSymbol)
	i++
	for i < *iterations {
		update(cellLine)
		printLine(cellLine, *aliveSymbol, *deadSymbol)
		i++
	}



}

func initialize(cells []rune) {
	seed := rand.NewSource(time.Now().UnixNano())
	randomiser := rand.New(seed)
	for i := range cells {
		state := randomiser.Intn(2)
		if state == 0 {
			cells[i] = deadStatus
		} else {
			cells[i] = aliveStatus
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
		cells[i] = transitionRules[pattern]
	}
}


func printLine(cells []rune, aliveSymbol, deadSymbol string) {
	for i,value := range cells {
		if value == deadStatus {
			fmt.Print(deadSymbol)
		} else if value == aliveStatus {
			fmt.Print(aliveSymbol)
		} else {
			log.Fatal(fmt.Sprintf("Unknown cell status %s at position %d", string(value), i+1))
			os.Exit(1)
		}
	}
	fmt.Println()
}


func createTransitionRules(ruleNumber int) map[string]rune {

	table := make(map[string]rune) // For testing purpose
	binaryString := toString(ruleNumber, 8)
	for i:= 7; i >= 0; i-- {
		pattern := toString(i, 3)
		state := binaryString[7-i]
		table[pattern] = rune(state)
	}
	return table

}

/* Returns the len string representation of an int as a binary*/
func toString(integer int, len int) string {
	return fmt.Sprintf("%0[1]*s",len,strconv.FormatInt(int64(integer),2))
}




