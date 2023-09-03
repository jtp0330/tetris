package game

import (
	"block"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// global vars
var game_grid = [][]string{}
var ROWS = 21
var COLS = 10
var ROW_END = 1
var shapes = []block.Block{}
var game_over = false
var windowWidth = 640
var windowHeight = 480
var block_on_grid = false

// block spawn point
var SPAWN = COLS / 2

// Player object
type User struct {
	userId   int
	userName string
	score    int
}
type Game struct {
	user      *User
	gameOver  bool
	currShape block.Block
}

// updates game state every frame
// inifintie loop may be obselte
func (g *Game) Update() error {

	//game loop
	//start first block drop

	if !block_on_grid {
		//spawn new block
		block_on_grid = true
		g.currShape = SpawnBlock(shapes)
	}
	//user input
	//TBD -> add quick down feature
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		for _, coord := range g.currShape {
			coord[0] = coord[0] - 1 //move left
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		for _, coord := range g.currShape {
			coord[0] = coord[0] + 1 //move right
		}
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		block.Rotate(g.currShape)
	}

	//move block down
	DropShape(g.currShape)

	//check for collision
	if Collision(g.currShape) {
		if GameOver() {
			g.gameOver = true
			return nil
		} else if IsRowFull() != 0 {
			//remove row and update score
			remove_row_index := IsRowFull()
			RemoveRow(game_grid, g.user, remove_row_index)
		} else {
			//spawn new block
			g.currShape = SpawnBlock(shapes)
		}
	}

	return nil
}

// draw current grid at each frame
func (g *Game) Draw(screen *ebiten.Image) {
	//draw background
	screen.Fill(color.Black)

	//draw score
	currScore := fmt.Sprintf("Score: %d", g.user.score)
	text.Draw(screen, currScore, nil, 10, 10, color.White)

	if g.gameOver {
		screen.Fill(color.Black)
		text.Draw(screen, "Game Over!", nil, 10, 10, color.White)
		text.Draw(screen, "{g.user.userName} score is {currScore}", nil, 10, 30, color.White)
	}

	//start game
	startGameLoop(g)

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

// move shape down until hits bottom or occupied space
// set coordinates of shape to occupied space when top condition is met
func DropShape(shape block.Block) {

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
func startGameLoop(g *Game) {

	fmt.Println("Welcome to Tetris!")
	fmt.Println("Please Enter a username before starting the game:")

	//create user and initalize game struct
	var username string
	fmt.Scan(&username)
	user := User{1, username, 0}
	g.user = &user
	g.gameOver = false

	fmt.Println("Game starting for user...: ", user.userName)
	time.Sleep(20)

	//create grid
	game_grid = CreateGrid()
	LoadBlocks()

	// //game loop
	// for playing_game {
	// 	//move block down
	// 	DropShape(currShape)
	// 	//check for collision
	// 	if Collision(currShape) {
	// 		if GameOver() {
	// 			g.gameOver = true
	// 			break
	// 		} else if IsRowFull() != 0 {
	// 			//remove row and update score
	// 			remove_row_index := IsRowFull()
	// 			RemoveRow(game_grid, &user, remove_row_index)
	// 		} else {
	// 			//spawn new block
	// 			currShape = SpawnBlock(shapes)
	// 		}
	// 	}
}

// start game
func CreateGame(game *Game) {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Tetris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
