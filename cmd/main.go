package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const colorReset = "\033[0m"

const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

type Player struct {
	name  string
	peice string
}

type GameState struct {
	playerOne     Player
	playerTwo     Player
	currentPlayer *Player
	moveCount     uint8
	board         Board
}

type Board [][3]string

func flattenBoard(b Board) []interface{} {
	var flatBoard []interface{}

	for _, cells := range b {
		for i := 0; i < len(cells); i++ {
			value := cells[i]

			if len(cells[i]) == 0 {
				value = " "
			}

			flatBoard = append(flatBoard, value)
		}
	}

	return flatBoard
}

func printBoard(b Board) {
	fb := flattenBoard(b)

	fmt.Println(fmt.Sprintf(`
		  A - B - C
		1 %s | %s | %s
		------------
		2 %s | %s | %s
		------------
		3 %s | %s | %s
	`, fb[0:]...))
}

func createPlayer(n int, ir *bufio.Reader) Player {
	fmt.Print(string(colorBlue), "Enter player one name: ", string(colorReset))
	name, err := ir.ReadString('\n')

	if err != nil {
		fmt.Println(string(colorRed), "Sorry an error ocurred while typing your name")
		log.Fatal("Error processing user input")
	}

	name = strings.TrimSuffix(name, "\n")

	peice := "o"
	if n == 2 {
		peice = "x"
	}

	player := Player{name, peice}
	return player
}

// missing arguments x int, y int
func winCheck() bool {
	won := true
	return won
}

func playerMove(gs GameState, ir *bufio.Reader) {
	winner := winCheck()
	if winner {
		message := fmt.Sprintf("Congratulations %s You won!", gs.currentPlayer.name)
		fmt.Println(string(colorCyan), message, string(colorReset))
		os.Exit(0)
	}
}

func gameLoop(gs GameState, ir *bufio.Reader) {
	playerMove(gs, ir)
	printBoard(gs.board)
	gameLoop(gs, ir)
}

func main() {
	var gameState GameState

	// Bootstrapping
	board := [][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	gameState.board = board
	gameState.moveCount = 0

	inputReader := bufio.NewReader(os.Stdin)

	for i := 1; i <= 2; i++ {
		p := createPlayer(i, inputReader)
		if i == 1 {
			gameState.playerOne = p
		} else {
			gameState.playerTwo = p
		}
	}

	gameState.currentPlayer = &gameState.playerOne

	// Star Game
	printBoard(board) // Print the inital board state
	gameLoop(gameState, inputReader)
}
