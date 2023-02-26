package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SudokuBoard [9][9]int



func main() {
	// Initialize a new Sudoku board
	board := newSudokuBoard()

	// Print the initial board
	fmt.Println("Initial board:")
	printBoard(board)

	// Prompt the user to select a difficulty level
	difficulty := promptDifficultyLevel()

	// Generate a Sudoku puzzle based on the selected difficulty level
	generatePuzzle(board, difficulty)

	// Print the puzzle
	fmt.Println("Puzzle:")
	printBoard(board)

	// Start the timer
	timer := time.NewTimer(90 * time.Second)

	// Allow the user to play the puzzle
	for i := 1; i <= 3; i++ {
		if solveBoard(board) {
			fmt.Println("Congratulations! You solved the puzzle.")
			return
		}
		if i == 3 {
			fmt.Println("You failed! Better luck next time.")
			return
		}
		fmt.Println("You have", 3-i, "attempts left.")
		fmt.Println("Would you like to solve the puzzle? (y/n)")
		var choice string
		fmt.Scanln(&choice)
		if choice == "y" {
			solveBoard(board)
			fmt.Println("Solution:")
			printBoard(board)
			return
		}
		<-timer.C
		timer.Reset(30 * time.Second)
		fmt.Println("Time's up! Would you like to solve the puzzle? (y/n)")
		fmt.Scanln(&choice)
		if choice == "y" {
			solveBoard(board)
			fmt.Println("Solution:")
			printBoard(board)
			return
		}
	}
	


}

func newSudokuBoard() {
	fmt.Print()
}



func solveBoard(board *SudokuBoard) bool {
	row, col := findEmptyCell(board)
	if row == -1 {
		return true // All cells have been filled
	}
	for num := 1; num <= 9; num++ {
		if isValidMove(board, row, col, num) {
			board[row][col] = num
			if solveBoard(board) {
				return true
			}
			board[row][col] = 0 // Backtrack if a solution cannot be found
		}
	}
	return false
}

// findEmptyCell finds the next empty cell on the board and returns its row and column.
// Returns -1 for both row and col if all cells have been filled.
func findEmptyCell(board *SudokuBoard) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

// isValidMove checks if a move is valid by checking the row, column, and 3x3 sub-grid for the given cell.
func isValidMove(board *SudokuBoard, row, col, num int) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// Check sub-grid
	subGridRow := (row / 3) * 3
	subGridCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[subGridRow+i][subGridCol+j] == num {
				return false
			}
		}
	}
	return true
}

// promptDifficultyLevel prompts the user to select a difficulty level and returns the selected level.
func promptDifficultyLevel() int {
	fmt.Println("Select a difficulty level:")
	fmt.Println("1 - Easy")
	fmt.Println("2 - Medium")
	fmt.Println("3 - Hard")
	var choice int
	fmt.Scanln(&choice)
	for choice < 1 || choice > 3 {
		fmt.Println("Invalid choice. Please select a valid option:")
		fmt.Scanln(&choice)
	}
	return choice
}

// generatePuzzle generates a Sudoku puzzle based on the selected difficulty level.
func generatePuzzle(board *SudokuBoard, difficulty int) {
	rand.Seed(time.Now().UnixNano())
	var numToRemove int
	switch difficulty {
	case 1:
		numToRemove = rand.Intn(20) + 30
	case 2:
		numToRemove = rand.Intn(20) + 45
	case 3:
		numToRemove = rand.Intn(20) + 60
	}
	for i := 0; i < numToRemove; i++ {
		row := rand.Intn(9)
		col := rand.Intn(9)
		if board[row][col] != 0 {
			board[row][col] = 0
		} else {
			i--
		}
	}
}

// printBoard prints the Sudoku board to the console.
func printBoard(board *SudokuBoard) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j], " ")
			if (j+1)%3 == 0 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 && i != 8 {
			fmt.Println("---------------------")
		}

	}
}
