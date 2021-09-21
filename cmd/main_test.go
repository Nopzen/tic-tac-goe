package main

import (
	"testing"

	"github.com/nopzen/tic-tac-goe/internal"
	"github.com/stretchr/testify/assert"
)

func TestWinCheck(t *testing.T) {
	p1 := Player{
		name:  "Test one",
		peice: "o",
		color: "fake",
	}

	p2 := Player{
		name:  "Test two",
		peice: "x",
		color: "fake",
	}

	b := internal.Board{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	gs := GameState{
		playerOne:     p1,
		playerTwo:     p2,
		currentPlayer: &p1,
		movesLeft:     9,
		board:         b,
	}

	result := winCheck(0, 0, gs)
	assert.Equal(t, false, result)

	// Should be a winning combination as all on
	// first row is, of current player peice
	gs.board[0][0] = gs.currentPlayer.peice
	gs.board[0][1] = gs.currentPlayer.peice
	gs.board[0][2] = gs.currentPlayer.peice

	result = winCheck(0, 0, gs)
	assert.Equal(t, true, result)

	// Should be a winning combination as all on
	// first column, is of current player peice
	gs.board[0][0] = gs.currentPlayer.peice
	gs.board[1][0] = gs.currentPlayer.peice
	gs.board[2][0] = gs.currentPlayer.peice

	result = winCheck(0, 0, gs)
	assert.Equal(t, true, result)

	// Should be a winning combination as all on
	// diagoal top left to bottom right, is of current player peice
	gs.board[0][0] = gs.currentPlayer.peice
	gs.board[1][1] = gs.currentPlayer.peice
	gs.board[2][2] = gs.currentPlayer.peice

	result = winCheck(0, 0, gs)
	assert.Equal(t, true, result)

	// Should be a winning combination as all on
	// diagoal top left to bottom right, is of current player peice
	gs.board[0][2] = gs.currentPlayer.peice
	gs.board[1][1] = gs.currentPlayer.peice
	gs.board[2][0] = gs.currentPlayer.peice

	result = winCheck(0, 0, gs)
	assert.Equal(t, true, result)
}
