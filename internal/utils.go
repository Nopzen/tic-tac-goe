package internal

import "fmt"

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

func PrintBoard(b Board) {
	fb := flattenBoard(b)

	fmt.Println(fmt.Sprintf(`
		  1 - 2 - 3
		A %s | %s | %s
		------------
		B %s | %s | %s
		------------
		C %s | %s | %s
	`, fb[0:]...))
}
