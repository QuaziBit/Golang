package main

import (
	"fmt"

	"github.com/go_projects/apps/tic_tac_toe/components"
)

// Player is structure in the player.go file
// basically Player used as label to call functions in the player.go file
var player components.Player
var ai components.AI

func initializePlayers(userName string) {

	if userName == "X" {
		player.SetPlayerInfo(1, userName, true)
		player.SetPlayerInfo(2, "O", false)
	} else {
		player.SetPlayerInfo(1, userName, true)
		player.SetPlayerInfo(2, "X", false)
	}

	// Init AI
	_, human, _, _ := player.GetPlayerInfo(true)
	_, name, _, _ := player.GetPlayerInfo(false)
	components.InitializeBoard()
	ai.InitializeAI(name, components.GetBoard(), human)
}

func options() string {
	fmt.Printf("\n")

	fmt.Printf("Enter [1] to be X or [2] to be O: ")
	var name string
	input := userInput()
	if input == 1 {
		name = "X"
	} else {
		name = "O"
	}
	return name
}

// userInput used to get input from user
// return type is int
func userInput() int {
	var input int

	// _ can be 0 or 1
	// 1: has correct input and 0: wrong input
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Printf("\nWrong input: %v\n\n", err)
	}

	return input
}

func showPlayers() {
	_, name1, _, _ := player.GetPlayerInfo(true)
	_, name2, _, _ := player.GetPlayerInfo(false)
	fmt.Printf("Player: %s vs. Player: %s\n", name1, name2)
}

func calculatePoints(board [9]string, humanName string, aiName string) {
	// Loop via game board and count X's and O's
	var tmp [3][3]string

	// Convert 1D array into 2D array
	n := 0
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[0]); j++ {
			tmp[i][j] = board[n]
			n++
		}
	}

	// Count X's and O's
	x := 0
	o := 0
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[0]); j++ {

			// Only horizontal check
			if tmp[i][j] == humanName {
				x++
			} else if tmp[i][j] == aiName {
				o++
			}
		}
	}

	fmt.Printf("\n=================================\n")
	fmt.Printf("Number of X: %d\n", x)
	fmt.Printf("Number of O: %d\n", o)

	// TEST
	if x == 3 {
		fmt.Printf("X won")
	} else if o == 3 {
		fmt.Printf("O won")
	} else {
		fmt.Printf("draw")
	}

	fmt.Printf("\n=================================\n")
}

func main() {

	// Store user input
	var input int

	fmt.Printf("Welcome to the Tic Tac Toe game!\n")

	// player can select its name [X] or [O]
	option := options()

	// create human player and AI player
	initializePlayers(option)

	// print players info
	showPlayers()

	// Print board
	components.PrintBoard()

	// start main game loop
	for i := 0; i < 9; i++ {

		// get human player info
		_, humanName, _, _ := player.GetPlayerInfo(true)

		// player will be asked to enter number
		fmt.Printf("Player [%s] enter index number or [-1] to stop game: ", humanName)
		input = userInput()

		// Stop game
		if input == -1 {
			break
		}

		// Human move
		components.UpdateBoard(input, humanName)

		// AI move
		ai.SetGameBoard(components.GetBoard())
		_, aiName, _, _ := player.GetPlayerInfo(false)
		components.UpdateBoard(ai.GetMove(), aiName)

		// calculate points
		calculatePoints(components.GetBoard(), humanName, aiName)

		// Show game board
		components.PrintBoard()
	}

}
