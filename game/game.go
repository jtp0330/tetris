package game

import (
	"fmt"
	"image/color"
	"log"
	"shape"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	ROWS         = 10
	COLS         = 21
	SPACING      = 15
	ROW_END      = 1
	windowWidth  = 640
	windowHeight = 480

	// block spawn point
	SPAWN_X = ROWS / 2
	SPAWN_Y = COLS / 2
)

// global vars
var game_grid = [][]string{}
var shapes = []shape.Shape{}
var block_on_grid = false
var game_over = false

// Player object
type User struct {
	userId   int
	userName string
	score    int
}
type Game struct {
	user      *User
	gameOver  bool
	currShape *shape.Shape
}

// updates game state every frame
// inifintie loop may be obselte
func (g *Game) Update() error {

	//game loop
	//start first block drop
	if !block_on_grid {
		//spawn new block
		block_on_grid = true
		g.currShape = SpawnShape(shapes)
	}
	//user input
	//TBD -> add quick down feature
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("Move left")
		g.currShape.MoveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("Move right")
		g.currShape.MoveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		shape.Rotate(g.currShape)
	}

	//move block down
	g.currShape.MoveDown()

	//check for collision
	if Collision(*g.currShape) {
		if GameOver() {
			g.gameOver = true
			return nil
		} else if IsRowFull() != 0 {
			//remove row and update score
			remove_row_index := IsRowFull()
			RemoveRow(game_grid, g.user, remove_row_index)
		} else {
			//spawn new block
			g.currShape = SpawnShape(shapes)
		}
	}

	return nil
}

// draw current grid at each frame
func (g *Game) Draw(screen *ebiten.Image) {
	//draw background
	screen.Fill(color.Black)
	//create font
	new_font := basicfont.Face7x13

	//render grid
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			text.Draw(screen, game_grid[i][j], new_font, windowWidth/4+(i*SPACING), windowHeight/4+(j*SPACING), color.RGBA{255, 255, 255, 255})
		}
	}

	//draw score
	currScore := fmt.Sprintf("Score: %d", g.user.score)
	text.Draw(screen, currScore, new_font, 5, windowHeight-10, color.White)

	if g.gameOver {
		screen.Fill(color.Black)
		text.Draw(screen, "Game Over!", new_font, windowWidth/2, windowHeight/2, color.White)
		text.Draw(screen, "{g.user.userName} score is {currScore}", new_font, windowWidth/2, windowHeight/2, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

// create initial grid
func CreateGrid() [][]string {
	grid := make([][]string, ROWS)
	//panic invalid pointer here
	for i := 0; i < ROWS; i++ {
		grid[i] = make([]string, COLS)
		for j := 0; j < COLS; j++ {
			//draw border
			if i < ROW_END {
				//first element
				grid[i][j] = "<!"
			} else if i >= ROWS-ROW_END {
				//last element
				grid[i][j] = "!>"
			} else {
				//draw empty space
				grid[i][j] = "."
			}
		}
	}
	//draw bottom border
	end := make([]string, COLS)
	for i := 0; i < COLS; i++ {
		end[i] = "<>"
	}
	print(grid)
	grid = append(grid, end)
	return grid
}

// creates and load shapes from shape.go
func LoadShapes(img *ebiten.Image, geoM *ebiten.GeoM) {

	var IShape = shape.Shape{shape.CreateBlock(img, geoM, 0, 1), shape.CreateBlock(img, geoM, 0, 2), shape.CreateBlock(img, geoM, 0, 3), shape.CreateBlock(img, geoM, 0, 4)} //I block
	var TShape = shape.Shape{shape.CreateBlock(img, geoM, 0, 0), shape.CreateBlock(img, geoM, 1, 0), shape.CreateBlock(img, geoM, 1, 1), shape.CreateBlock(img, geoM, 2, 0)} //T block
	var BShape = shape.Shape{shape.CreateBlock(img, geoM, 0, 0), shape.CreateBlock(img, geoM, 1, 0), shape.CreateBlock(img, geoM, 0, 1), shape.CreateBlock(img, geoM, 1, 1)} //Square block
	var SShape = shape.Shape{shape.CreateBlock(img, geoM, 0, 0), shape.CreateBlock(img, geoM, 1, 0), shape.CreateBlock(img, geoM, 1, 1), shape.CreateBlock(img, geoM, 2, 1)} //S block

	//load shapes defined in shape.go
	shapes = append(shapes, IShape)
	shapes = append(shapes, BShape)
	shapes = append(shapes, SShape)
	shapes = append(shapes, TShape)
}

// create block randomly in 4 different ways:
func SpawnShape(shapes []shape.Shape) *shape.Shape {
	//block shift left with an XOR operation
	//block shift down with an AND operation
	shape_index := 0 //randomize this
	shape_to_be_dropped := shapes[shape_index]
	return &shape_to_be_dropped
}

// modify function so that it accounts for checking rows only below it for collision
// if next coordinates down hit a '[]', then collision is true
// returns true if shape hits bottom or occupied space
func Collision(currentShape shape.Shape) bool {
	//check if shape hits bottom or occupied space

	for _, block := range currentShape {
		//check if shape hits bottom or occupied space

		if block.Point_x == ROWS-1 || game_grid[block.Point_y+1][block.Point_y] == "[]" {
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
	//genearte blocks and grid
	LoadShapes(&ebiten.Image{}, &ebiten.GeoM{})
	game_grid = CreateGrid()
}

// start game
func CreateGame(game *Game) {
	ebiten.SetWindowSize(windowWidth*2, windowHeight*2)
	ebiten.SetWindowTitle("Tetris")
	startGameLoop(game)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
