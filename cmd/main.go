package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nopzen/tic-tac-goe/internal"
)

const colorReset = "\033[0m"

const colorError = "\033[31m"
const colorSystem = "\033[36m"
const colorP1 = "\033[32m"
const colorP2 = "\033[35m"

type Player struct {
	name  string
	peice string
	color string
}

type GameState struct {
	playerOne     Player
	playerTwo     Player
	currentPlayer *Player
	movesLeft     int
	board         internal.Board
}

func createPlayer(n int, ir *bufio.Reader) Player {
	fmt.Print(string(colorSystem), fmt.Sprintf("Enter player %d name: ", n), string(colorReset))
	name, err := ir.ReadString('\n')

	if err != nil {
		fmt.Println(string(colorError), "Sorry an error ocurred while typing your name")
		log.Fatal("Error processing user input")
	}

	name = strings.TrimSuffix(name, "\n")

	peice := "o"
	color := colorP1
	if n == 2 {
		peice = "x"
		color = colorP2
	}

	player := Player{name, peice, color}
	return player
}

func winCheck(x int, y int, gs GameState) bool {
	won := false
	boardWidth := len(gs.board[x])

	// Check rows
	for i := 0; i < boardWidth; i++ {
		if gs.board[x][i] != gs.currentPlayer.peice {
			break
		}

		if i == boardWidth-1 {
			won = true
		}
	}

	// check columns
	for i := 0; i < boardWidth; i++ {
		if gs.board[i][y] != gs.currentPlayer.peice {
			break
		}

		if i == boardWidth-1 {
			won = true
		}
	}

	// check a1 to c3 diagonal
	if x == y {
		for i := 0; i < boardWidth; i++ {
			if gs.board[i][i] != gs.currentPlayer.peice {
				break
			}

			if i == boardWidth-1 {
				won = true
			}
		}
	}

	// check a3 to c1 reversed tiagonal
	if x+y == boardWidth-1 {
		for i := 0; i < boardWidth; i++ {
			if gs.board[i][(boardWidth-1)-i] != gs.currentPlayer.peice {
				break
			}

			if i == boardWidth-1 {
				won = true
			}
		}

	}

	return won
}

func playerMove(gs GameState, ir *bufio.Reader) error {

	fmt.Printf("%s%s it's your turn to move: %s", string(gs.currentPlayer.color), gs.currentPlayer.name, string(colorReset))
	cordinate, err := ir.ReadString('\n')
	if err != nil {
		fmt.Println(string(colorError), "Sorry an error while you tried to make your move.", string(colorReset))
		return errors.New("could not read user cordinate input")
	}

	cordinate = strings.TrimSuffix(cordinate, "\n")

	var x int
	var y int

	if len(cordinate) != 2 {
		return errors.New("please enter a valid cordinate eg. b2")
	}

	switch string(cordinate[0]) {
	case "a":
		x = 0
	case "b":
		x = 1
	case "c":
		x = 2
	default:
		fmt.Println(string(colorError), fmt.Sprintf("given row (%s) is not valid try again, should be (a-c)", string(cordinate[0])), string(colorReset))
		return errors.New("failed to process X cordinate")
	}

	cy, err := strconv.Atoi(string(cordinate[1]))
	if err != nil {
		fmt.Println(string(colorError), fmt.Sprintf("Given colum (%d) is not valid try again, should be (1-3)", cy), string(colorReset))
		return errors.New("failed to process Y cordinate")
	}

	// subtract 1 to adhere to 0 index based arrays
	y = cy - 1

	if len(gs.board[x][y]) != 0 {
		fmt.Println(string(colorError), fmt.Sprintf("Field is already taken by '%s', try again.", gs.board[x][y]), string(colorReset))
		return errors.New("failed to process cordinate, field already taken")
	}

	// insert player peice at cordinates
	gs.board[x][y] = gs.currentPlayer.peice

	internal.PrintBoard(gs.board)

	winner := winCheck(x, y, gs)
	if winner {
		message := fmt.Sprintf("Congratulations %s You won!", gs.currentPlayer.name)
		fmt.Println(string(colorSystem), message, string(colorReset))
		os.Exit(0)
	}

	if gs.movesLeft == 0 {
		fmt.Println(string(colorSystem), "Tie game! better luck next time.", string(colorReset))
		os.Exit(0)
	}

	return nil
}

func gameLoop(gs GameState, ir *bufio.Reader) {
	err := playerMove(gs, ir)
	if err != nil {
		fmt.Println(string(colorError), err)
		gameLoop(gs, ir)
		return
	}

	if gs.currentPlayer.peice == gs.playerOne.peice {
		gs.currentPlayer = &gs.playerTwo
	} else {
		gs.currentPlayer = &gs.playerOne
	}

	gs.movesLeft = gs.movesLeft - 1
	gameLoop(gs, ir)
}

func main() {
	var gameState GameState
	gameState.movesLeft = 0

	// Bootstrapping
	board := [][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	gameState.board = board
	for _, cols := range gameState.board {
		gameState.movesLeft += len(cols)
	}

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
	internal.PrintBoard(board) // Print the inital board state
	gameState.movesLeft = gameState.movesLeft - 1

	for {
		gameLoop(gameState, inputReader)
	}
}
