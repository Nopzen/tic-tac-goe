package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenBoard(t *testing.T) {
	board := [][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	expected := []interface{}{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	result := flattenBoard(board)

	assert.Equal(t, expected, result)
}
