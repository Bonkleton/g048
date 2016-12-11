package main

import (
	"fmt"
	"os"
	"strconv"
)

// master game function
func game(size, base, pieces int) int {
	fmt.Println(" - Welcome to g048 by Bonkleton! -")

	// board initialized, random tiles placed on board
	gb := rtg(rcg(pieces, size), base, gbInit(make([][]string, size)))
	printGame(gb)

	// master game loop
	for {
		// see if the player has lost yet
		if gameOver(gb) {
			if continueGame() {
				return 1
			} else {
				return 0
			}
		}

		// take input
		var input string
		fmt.Scanln(&input)

		// decide based on input
		switch input {
		case "x":
			saveGame(gb) // save gameboard to text file for later
			return 0
		case "s":
			gb = rtg(rcg(1, size), base, crunchAll(size, "d", gb)) // add random tile to crunched gameboard
			printGame(gb)
		case "w":
			gb = rtg(rcg(1, size), base, crunchAll(size, "u", gb)) // add random tile to crunched gameboard
			printGame(gb)
		case "d":
			gb = rtg(rcg(1, size), base, crunchAll(size, "r", gb)) // add random tile to crunched gameboard
			printGame(gb)
		case "a":
			gb = rtg(rcg(1, size), base, crunchAll(size, "l", gb)) // add random tile to crunched gameboard
			printGame(gb)
		case "42":
			fmt.Println("42 is the answer, but a more powerful computer than I must find the question!")
		case "69":
			fmt.Println("lol 69")
		case "420":
			fmt.Println("dank meth is lit fam")
		default:
			fmt.Println("Invalid!")
		}
	}
	return 0
}

func main() {

	var size, base, pieces int
	a := append(os.Args[1:], "", "", "") // get cmd line args

	if a[0] == "" { // handle size arg
		size = 4
	} else if size < 4 {
		fmt.Println("Gameboard too small! Defaulting to size 4...")
		size = 4
	} else {
		size, _ = strconv.Atoi(a[0])
	}

	if a[1] == "" { // handle base arg
		base = 2
	} else if base == 0 {
		fmt.Println("Infinite game impossible! Defaulting to base 2...")
		base, _ = strconv.Atoi(a[1])
	} else {
		base, _ = strconv.Atoi(a[1])
	}

	if a[2] == "" { // handle pieces arg
		pieces = 2
	} else if pieces == 0 {
		fmt.Println("Empty starting board impossible! Defaulting to 2 pieces...")
		pieces, _ = strconv.Atoi(a[2])
	} else {
		pieces, _ = strconv.Atoi(a[2])
	}

	// start game
	for {
		if game(size, base, pieces) == 0 {
			break
		}
	}
}
