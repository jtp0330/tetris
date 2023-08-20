package game

import (
	"block"
	"fmt"
	"grid"
	"math/rand"
	"time"
)

// global vars
var game_grid = [][]string{}
var ROWS = 21
var COLS = 10
var ROW_END = 1
var shapes = []block.Block{}
var game_over = false
var playing_game = false

// block spawn point
var SPAWN = COLS / 2

// Player object
type User struct {
	userId   int
	userName string
	score    int
}

func LoadBlocks() {
	//load shapes defined in block.go
	shapes = append(shapes, block.IBlock)
	shapes = append(shapes, block.BBlock)
	shapes = append(shapes, block.SBlock)
	shapes = append(shapes, block.TBlock)
}

// create block randomly in 4 different ways:
func SpawnBlock(shapes []block.Block) block.Block {
	//block shift left with an XOR operation
	//block shift down with an AND operation
	block_index := rand.New(rand.NewSource(4)).Int()
	block_to_be_dropped := shapes[block_index]
	//by def, using range on an array returns the index and the value,
	//we only care about value, thus ignore index with _
	for _, coord := range block_to_be_dropped {
		coord_x := coord[0]
		coord_y := coord[1]
		game_grid[coord_x+SPAWN][coord_y+SPAWN] = "[]"

	}
	return block_to_be_dropped
}

func MoveShape(shape block.Block) {
	//move shape down until hits bottom or occupied space
	//set coordinates of shape to occupied space when top condition is met
	for _, coord := range shape {
		coord[0], coord[1] = coord[0]+1, coord[1]+1 //move down
	}
}

// returns true if shape hits bottom or occupied space
func Collision(currentShape block.Block) bool {
	//check if shape hits bottom or occupied space
	for _, coord := range currentShape {
		//check if shape hits bottom or occupied space
		if coord[0] == 1 || coord[1] == 1 || coord[0] == ROWS-1 || coord[1] == COLS-1 {
			return true
		}
	}
	return false
}

func IsRowFull() bool {
	//check if there are any full rows
	//if there are, remove them and move everything down
	//update score
	isFull := false
	for _, rows := range game_grid {
		for j := 0; j < COLS; j++ {
			if rows[j] != "[]" {
				isFull = true
			}
		}
	}
	return isFull
}

func GameOver() bool {
	//check if game is over
	//if game is over, update score and exit game loop
	//check top row, and if any of the columns are occupied, and not a full row, game over
	for i := 0; i < COLS; i++ {
		if game_grid[0][i] != "." && !IsRowFull() {
			game_over = true
		}
	}
	return game_over
}

func RemoveRow(grid [][]string, user User) {
	//remove row
	//move everything down
	//update score
	user.score += 10
}

// game logic
func startGameLoop() {
	//create user

	fmt.Println("Welcome to Tetris!")
	fmt.Println("Please Enter a username before starting the game:")

	var username string
	fmt.Scan(&username)
	//create user
	user := User{1, username, 0}
	fmt.Println("Game starting for user...: ", user.userName)
	time.Sleep(5)

	//create grid
	game_grid = grid.CreateGrid()
	playing_game = true

	//game loop
	for playing_game {
		//spawn block
		currShape := SpawnBlock(shapes)
		//move block down
		MoveShape(currShape)
		//check for collision
		if Collision(currShape) {
			if GameOver() {
				break
			} else if IsRowFull() {
				//remove full row
				//move everything down
				//update score
				RemoveRow(game_grid, user)
			}
		}
	}
	//game over message
	fmt.Println("Game over!")
	fmt.Println("Your Final score is: ", user.score)
}

func CreateGame() {
	LoadBlocks()
	startGameLoop()
}
