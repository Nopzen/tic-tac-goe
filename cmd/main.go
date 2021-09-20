package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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
	color string
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
	color := colorGreen
	if n == 2 {
		peice = "x"
		color = colorPurple
	}

	player := Player{name, peice, color}
	return player
}

func winCheck(x int, y int) bool {
	won := false
	return won
}

func playerMove(gs GameState, ir *bufio.Reader) error {

	fmt.Printf("%s%s it's your turn to move: %s", string(gs.currentPlayer.color), gs.currentPlayer.name, string(colorReset))
	cordinate, err := ir.ReadString('\n')
	if err != nil {
		fmt.Println(string(colorRed), "Sorry an error while you tried to make your move.", string(colorReset))
		return errors.New("Could not read user cordinate input")
	}

	var x int
	var y int

	switch string(cordinate[0]) {
	case "a":
		x = 0
	case "b":
		x = 1
	case "c":
		x = 2
	default:
		// should retry
		return errors.New("Failed to parse user given cordinate X")
	}

	cordinateY, err := strconv.Atoi(string(cordinate[1]))
	if err != nil {
		// Should retry
		return errors.New("Failed to parse user given cordinate Y")
	}

	// subtract 1 to adhere to 0 index based arrays
	y = cordinateY - 1

	// insert player peice at cordinates
	gs.board[x][y] = gs.currentPlayer.peice

	winner := winCheck(x, y)
	if winner {
		message := fmt.Sprintf("Congratulations %s You won!", gs.currentPlayer.name)
		fmt.Println(string(colorCyan), message, string(colorReset))
		os.Exit(0)
	}

	return nil
}

func gameLoop(gs GameState, ir *bufio.Reader) {
	err := playerMove(gs, ir)
	if err != nil {
		log.Fatal(err)
	}

	printBoard(gs.board)

	if gs.currentPlayer.peice == gs.playerOne.peice {
		gs.currentPlayer = &gs.playerTwo
	} else {
		gs.currentPlayer = &gs.playerOne
	}

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
