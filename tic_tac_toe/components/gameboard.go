package components

import (
	"fmt"
)

var isInitialized bool
var board [9]string

// InitializeBoard build game board
func InitializeBoard() {
	for i := 0; i < len(board); i++ {
		n := fmt.Sprintf("%d", (i + 1)) // converting int into string
		board[i] = n
	}
	isInitialized = true
}

// UpdateBoard updates gameboard
// Arguments: index int
func UpdateBoard(index int, symbol string) {
	index--
	board[index] = symbol
}

// PrintBoard prints gameboard
func PrintBoard() {

	// At the beginning of game the initializeBoard() method will be called
	if !isInitialized {
		InitializeBoard()
	}

	fmt.Printf("\n")
	n := 0
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s ", board[i])

		// used to print new line after [3] elements of the board
		if n == 2 {
			n = 0
			n--
			fmt.Printf("\n")
		}
		n++
	}
	fmt.Printf("\n")
}

// GetBoard get current game board
func GetBoard() [9]string {
	return board
}
