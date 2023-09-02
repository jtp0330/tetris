package game

import (
	"block"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// global vars
var game_grid = [][]string{}
var ROWS = 21
var COLS = 10
var ROW_END = 1
var shapes = []block.Block{}
var game_over = false
var playing_game = false
var windowWidth = 640
var windowHeight = 480

// block spawn point
var SPAWN = COLS / 2

// Player object
type User struct {
	userId   int
	userName string
	score    int
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

// draw current grid at each frame
func (g *Game) Draw(screen *ebiten.Image) {

	//draw grid
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			//draw grid
			ebitenutil.DebugPrintAt(screen, game_grid[i][j], i, j)
		}
		fmt.Println()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// create initial grid
func CreateGrid() [][]string {
	grid := make([][]string, ROWS)
	for i := 0; i < ROWS; i++ {
		grid[i] = make([]string, COLS)
		for j := 0; j < COLS; j++ {
			//draw border
			if j < ROW_END {
				//first element
				grid[i][j] = "<!"
			} else if j >= COLS-ROW_END {
				//last element
				grid[i][j] = "!>"
			} else {
				//draw empty space
				grid[i][j] = " ."
			}
		}
	}
	//draw bottom border
	end := make([]string, COLS)
	for i := 0; i < COLS; i++ {
		end[i] = "<>"
	}
	grid[ROWS-1] = end

	fmt.Println(grid)
	return grid
}

// load blocks
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

// modify function so that it accounts for checking rows only below it for collision
// if next coordinates down hit a '[]', then collision is true
// returns true if shape hits bottom or occupied space
func Collision(currentShape block.Block) bool {
	//check if shape hits bottom or occupied space
	for _, coord := range currentShape {
		//check if shape hits bottom or occupied space
		if game_grid[coord[0]][coord[1]+1] == "[]" || coord[1] == COLS-1 {
			return true
		}
	}
	return false
}

func IsRowFull() int {
	//check if there are any full rows
	//if there are, remove them and move everything down
	//update score
	row_index := 0
	for i, row := range game_grid {
		for j := 0; j < COLS; j++ {
			if row[j] != "[]" {
				row_index = i
				break
			}
		}
	}
	return row_index
}

func GameOver() bool {
	//check if game is over
	//if game is over, update score and exit game loop
	//check top row, and if any of the columns are occupied, and not a full row, game over
	for i := 0; i < COLS; i++ {
		if game_grid[0][i] != "." && IsRowFull() == 0 {
			game_over = true
		}
	}
	return game_over
}

func RemoveRow(grid [][]string, user *User, row_index int) {
	//remove row by moving everything less than row_index down
	for i := 0; i < row_index; i++ {
		grid[i] = grid[i+1]
	}
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
	game_grid = CreateGrid()
	LoadBlocks()
	//start first block drop
	currShape := SpawnBlock(shapes)
	playing_game = true

	//game loop
	for playing_game {
		//move block down
		MoveShape(currShape)
		//check for collision
		if Collision(currShape) {
			if GameOver() {
				break
			} else if IsRowFull() != 0 {
				//remove row and update score
				remove_row_index := IsRowFull()
				RemoveRow(game_grid, &user, remove_row_index)
			} else {
				//spawn new block
				currShape = SpawnBlock(shapes)
			}
		}
	}
	//game over message
	fmt.Println("Game over!")
	fmt.Println("Your Final score is: ", user.score)
}

// start game
func CreateGame() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Tetris")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	//start game
	startGameLoop()
}
