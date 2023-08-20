package block

type Coordinates [2]int
type Block [4]Coordinates

// block patterns
var IBlock = Block{{0, 1}, {0, 2}, {0, 3}, {0, 4}} //I block
var TBlock = Block{{1, 0}, {1, 1}, {2, 0}, {0, 1}} //T block
var BBlock = Block{{0, 0}, {1, 0}, {0, 1}, {1, 1}} //Square block
var SBlock = Block{{0, 0}, {1, 0}, {1, 1}, {2, 1}} //S block

// create block randomly in 4 different ways:
func Rotate(block Block) (rotated_block Block) {
	//block shift left with an XOR operation
	//block shift right with an OR operation
	//block shift down with an AND operation
	//block shift up with an AND NOT operation
	for coord := range block {
		//block[i] = block[i] << 1
		coord ^= 1
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
