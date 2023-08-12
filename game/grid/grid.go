package grid

import "fmt"

// global vars
var ROWS = 21
var COLS = 10
var ROW_END = 1

func CreateGrid() {
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
}
