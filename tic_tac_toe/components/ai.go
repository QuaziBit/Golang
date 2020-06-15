package components

import "fmt"

// AI structure
type AI struct {
	name      string
	level     int
	board     [9]string
	move      int
	humanName string
}

var aiBot AI

func generateInput(col int, j int) int {
	aiInput := -1
	if col == 0 && j == 0 {
		aiInput = 1
	} else if col == 0 && j == 1 {
		aiInput = 2
	} else if col == 0 && j == 2 {
		aiInput = 3
	} else if col == 1 && j == 0 {
		aiInput = 4
	} else if col == 1 && j == 1 {
		aiInput = 5
	} else if col == 1 && j == 2 {
		aiInput = 6
	} else if col == 2 && j == 0 {
		aiInput = 7
	} else if col == 2 && j == 1 {
		aiInput = 8
	} else if col == 2 && j == 2 {
		aiInput = 9
	}

	return aiInput
}

// bestMove function used to calculate best AI move
// Arguments: col int => column, row int => row, board 2D array of the game board
func bestMove(col int, row int, board [3][3]string) int {

	aiInput := -1

	// loop via specific game board row and insert X or O basically X or O is AI name
	for j := 0; j < len(board[0]); j++ {
		if board[col][j] != aiBot.humanName && board[col][j] != aiBot.name {
			board[col][j] = aiBot.name

			// Generate AI input
			aiInput = generateInput(col, j)
		}
	}

	// TEST
	//==================================================================================//
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	//==================================================================================//

	return aiInput
}

func aiMove() {
	// analyze game board and find the beast movie option

	var tmp [3][3]string
	tmpBoard := aiBot.board

	// 1 2 3 - 4 5 6 - 7 8 9
	// 1|2|3
	// 4|5|6
	// 7|8|9

	// 0.0|0.1|0.2
	// 1.0|1.1|1.2
	// 2.0|2.1|2.2

	// Min-Max
	// X|O|O =>  0
	// O|X|X => -1
	// O|O|O => +1

	// Convert 1D array into 2D array
	n := 0
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[0]); j++ {
			tmp[i][j] = tmpBoard[n]
			n++
		}
	}

	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[0]); j++ {

			// Only Horizontal loop
			if tmp[i][j] == aiBot.humanName {
				// pass the human symbol position, and [tmp] game board
				aiBot.move = bestMove(i, j, tmp)
			}

		}
	}

	// TEST
	//aiBot.move = 5
}

// InitializeAI creates AI bot
func (a AI) InitializeAI(aiName string, gameBoard [9]string, humanName string) {
	aiBot = AI{aiName, 1, gameBoard, -1, humanName}
}

// SetGameBoard store current game board
func (a AI) SetGameBoard(board [9]string) {
	aiBot.board = board
}

// GetMove return AI move
func (a AI) GetMove() (index int) {
	aiMove()
	return aiBot.move
}

/*
	1|2|3
	4|5|6
	7|8|9

	player: X

	1|X|3 [0]
	4|5|6 [1]
	7|8|9 [2]

	Take [0] and loop via 1|X|3, AI will put O => O|X|O
	if total point is 0 that it is best move

*/
