package grid

// //package redudant as same as game.go
// import (
// 	"fmt"

// 	"github.com/hajimehoshi/ebiten/ebitenutil"
// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type Game struct{}

// // global vars
// var ROWS = 21
// var COLS = 10
// var ROW_END = 1

// // create initial grid
// func CreateGrid() [][]string {
// 	grid := make([][]string, ROWS)
// 	for i := 0; i < ROWS; i++ {
// 		grid[i] = make([]string, COLS)
// 		for j := 0; j < COLS; j++ {
// 			//draw border
// 			if j < ROW_END {
// 				//first element
// 				grid[i][j] = "<!"
// 			} else if j >= COLS-ROW_END {
// 				//last element
// 				grid[i][j] = "!>"
// 			} else {
// 				//draw empty space
// 				grid[i][j] = " ."
// 			}
// 		}
// 	}
// 	//draw bottom border
// 	end := make([]string, COLS)
// 	for i := 0; i < COLS; i++ {
// 		end[i] = "<>"
// 	}
// 	grid[ROWS-1] = end

// 	fmt.Println(grid)
// 	return grid
// }

// func (g *Game) Update() error {
// 	return nil
// }

// // draw current grid at each frame
// func (g *Game) DrawGrid(grid [][]string, screen *ebiten.Image) {
// 	//draw grid
// 	for i := 0; i < ROWS; i++ {
// 		for j := 0; j < COLS; j++ {
// 			//draw grid
// 			ebitenutil.DebugPrint(screen, grid[i][j])
// 		}
// 		fmt.Println()
// 	}
// }
// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 320, 240
// }
