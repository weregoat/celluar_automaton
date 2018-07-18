package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
	"strconv"
	"os"
	"log"
	"strings"
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
	seed := flag.String("seed", "", "String representing the first cell line")
	flag.Parse()
	if ruleNumber < 0 || ruleNumber > 255 {
		log.Fatal(fmt.Sprintf("Invalid rule number %d", ruleNumber))
		os.Exit(1)
	}
	transitionRules = createTransitionRules(ruleNumber)
	cellLine := make([]rune, *cells, *cells)
	if len(*seed) > 0 {
		parseSeed(*seed, *deadSymbol, *aliveSymbol, *cells, cellLine)
	} else {
		initialize(cellLine)
	}
	i := 0
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

func parseSeed(seed, dead, live string, length int, line []rune) {
	/* Check that no alive and dead symbols align with the seed */
	/* Replace all dead or alive symbols within the string */
	emptyLine := strings.Replace(strings.Replace(seed, live, "", -1), dead, "", -1)
	if len(emptyLine) > 0 {
		log.Fatal(fmt.Sprintf("Invalid symbols in seed: %s", emptyLine))
		os.Exit(1)
	}
	if len(seed) > length {
		log.Fatal(fmt.Sprintf("Seed line is too long"))
		os.Exit(1)
	}
	/* Pad the string to length */
	padding := strings.Repeat(dead, length - len(seed))
	cellLine := fmt.Sprintf("%s%s",padding, seed)
	/* Insert the elements into the rune array */
	for i,v := range cellLine {
		if string(v) == dead {
			line[i] = deadStatus
		} else {
			line[i] = aliveStatus
		}
	}

}




