package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// initializes gameboard
func gbInit(gb [][]string) [][]string {
	m := len(gb)
	for i := 0; i < m; i++ { // set all blanks
		gb[i] = make([]string, m)
		for j := 0; j < m; j++ {
			gb[i][j] = "-"
		}
	}
	return gb
}

// prints the gameboard to terminal
func printGame(gb [][]string) {
	// print top border
	fmt.Printf("+-")
	for i := 0; i < len(gb); i++ {
		fmt.Printf("--")
	}
	fmt.Printf("+")
	fmt.Println()

	// print the board content
	for i := 0; i < len(gb); i++ {
		fmt.Printf("| ")
		for j := 0; j < len(gb); j++ {
			fmt.Print(gb[i][j])
			fmt.Print(" ")
		}
		fmt.Printf("|")
		fmt.Println()
	}

	// print bottom border
	fmt.Printf("+-")
	for i := 0; i < len(gb); i++ {
		fmt.Printf("--")
	}
	fmt.Printf("+")
	fmt.Println()

	// print controls/instructions
	fmt.Println("Choose direction to swipe the board:")
	fmt.Print("   x: exit \n   s: down \n   w: up \n   a: right\n   d: left \n>>> ")
}

// random coordinate generator
func rcg(p, m int) []int {
	coords := make([]int, p*2) // initialize list of coords to test
	for i := 1; i < len(coords); i++ {
		coords[i-1] = rand.Intn(m)
		coords[i] = rand.Intn(m) // populate list with random numbers
	}
	return coords
}

// takes list of coords and generates random tiles
func rtg(c []int, b int, gb [][]string) [][]string {
	i := 1
	for i < len(c) {
		if gb[c[i]][c[i-1]] != "-" {
			d := rcg(1, len(gb))
			c[i] = d[0]
			c[i-1] = d[1]
			continue
		} else {
			gb[c[i]][c[i-1]] = strconv.Itoa(maybeDbl(b))
			i = i + 2
		}
	}
	return gb
}

// doubles its input 1/10 of the time
func maybeDbl(x int) int {
	if rand.Intn(10) == 0 {
		return x * 2
	} else {
		return x
	}
}

// loops through the rows on the board, crunching them
func crunchAll(m int, s string, gb [][]string) [][]string {
	for i := 0; i < m; i++ {
		v, crunched := crunch(exportSlice(i, m, gb, s))
		if crunched {
			importSlice(i, v, gb, s)
		} else {
			continue
		}
	}
	return gb
}

// crunches a row of the board in direction swiped
func crunch(v []string) ([]string, bool) {
	res := false
	for i := 1; i < len(v); i++ {
		for j := len(v) - 1; j > 0; j-- {
			if v[j] == "-" { // if there's an empty tile in the crushing direction
				res = true
				v[j] = v[j-1] // close empty gap
				v[j-1] = "-"
			} else if v[j] == v[j-1] { //if there's two matching tiles in the crunching direction
				res = true
				a, _ := strconv.Atoi(v[j]) // get int value from v[i]
				v[j] = strconv.Itoa(a * 2) // combine tiles
				v[j-1] = "-"
			}
		}
	}
	return v, res
}

// gets slices from main game board for easy crunching
func exportSlice(k, m int, board [][]string, s string) []string {
	v := make([]string, m) // declare new slice
	switch s {
	case "d":
		for i := 0; i < m; i++ { // for as many elements as rows, increasing
			v[i] = board[i][k] // fill in new slice
		}
	case "u":
		for i := m; i > 0; i-- { // for as many elements as rows, decreasing
			v[m-i] = board[i-1][k] // fill in new slice
		}
	case "r":
		for i := 0; i < m; i++ { // for as many elements as columns, increasing
			v[i] = board[k][i] // fill in new slice
		}
	case "l":
		for i := m; i > 0; i-- { // for as many elements as rcolumns, decreasing
			v[m-i] = board[k][i-1] // fill in new slice
		}
	}
	return v
}

// updates slices in main game board after they've been crunched
// same logic as exportSlice but with inverted assignment statements
func importSlice(k int, v []string, board [][]string, s string) []string {
	m := len(v) // get length of v from exportSlice
	switch s {
	case "d":
		for i := 0; i < m; i++ { // for as many elements as rows, increasing
			board[i][k] = v[i] // fill in board
		}
	case "u":
		for i := m; i > 0; i-- { // for as many elements as rows, decreasing
			board[i-1][k] = v[m-i] // fill in board
		}
	case "r":
		for i := 0; i < m; i++ { // for as many elements as columns, increasing
			board[k][i] = v[i] // fill in board
		}
	case "l":
		for i := m; i > 0; i-- { // for as many elements as columns, decreasing
			board[k][i-1] = v[m-i] // fill in board
		}
	}
	return v
}

func gameOver(gb [][]string) bool {
	res := true
	for i := 0; i < len(gb); i++ {
		for j := 0; j < len(gb); j++ {
			if gb[i][j] == "-" {
				res = false
			} else {
				continue
			}
		}
	}
	return res
}

func continueGame() bool {
	fmt.Println("Game over!")
	fmt.Println("Play again?\n    yes: continue\n   no: exit")

	// take input
	var input string
	fmt.Scanln(&input)
	switch input {
	case "yes":
		return true
	case "no":
		return false
	default:
		fmt.Println("Invalid!")
		return continueGame() // call self to loop if invalid input
	}
}

func saveGame(gb [][]string) {
	fmt.Println("Game saved!")
}
