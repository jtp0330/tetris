package block

import (
	"tetris/game/grid"
)

type Coordinates [2]int

// block patterns
var Block1 = [4]Coordinates{{0, 1}, {0, 2}, {0, 3}, {0, 4}} //I block
var Block2 = [4]Coordinates{{1, 0}, {1, 1}, {2, 0}, {3, 0}} //T block
var Block3 = [4]Coordinates{{1, 0}, {1, 0}, {0, 1}, {1, 1}} //Z block
var Block4 = [4]Coordinates{{1, 0}, {0, 1}, {1, 1}, {2, 0}} //S block
// block spawn point
var SPAWN = grid.COLS / 2

// create block randomly in 4 different ways:
func Rotate(block [4]Coordinates) (rotated_block [4]Coordinates) {
	//block shift left with an XOR operation
	//block shift right with an OR operation
	//block shift down with an AND operation
	//block shift up with an AND NOT operation
	for i := 0; i < len(block); i++ {
		//block[i] = block[i] << 1
		for j := 0; j < len(block); j++ {
			block[i][j] = block[i][j] ^ 1
		}
	}

	return
}

// create block randomly in 4 different ways:
func Quick_Down(grid [][]string, block [4]Coordinates) (rotated_block [4]Coordinates) {
	//block shift left with an XOR operation
	//block shift down with an AND operation
	for i := 0; i < len(block); i++ {
		//block[i] = block[i] << 1
		for j := 0; j < len(block); j++ {
			block[i][j] = block[i][j] ^ 1
		}
	}

	return
}
