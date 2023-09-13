package shape

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// bug fixed, golang field names must be capitalized to be exported
type Block struct {
	Point_x int
	Point_y int
	Img     *ebiten.Image
	GeoM    *ebiten.GeoM
}

type Shape [4]*Block

func CreateBlock(img *ebiten.Image, geoM *ebiten.GeoM, tx int, ty int) *Block {
	//create block image
	blockImage := Block{tx, ty, img, geoM}
	return &blockImage
}

// create block randomly in 4 different ways:
func Rotate(shape *Shape) *Shape {
	//block shift left with an XOR operation
	//block shift right with an OR operation
	//block shift down with an AND operation
	//block shift up with an AND NOT operation
	for _, block := range *shape {
		block.Point_x, block.Point_y = block.Point_x^1, block.Point_y^1
	}

	return shape
}

func Quick_Down(grid [][]string, block [4]Block) {
	//block shift left with an XOR operation
	//block shift down with an AND operation
	fmt.Println("TBD")
}

func (shape Shape) MoveDown() {
	for _, block := range shape {
		block.GeoM.Translate(0, 1.0)
	}
}

func (shape Shape) MoveLeft() {
	for _, block := range shape {
		block.GeoM.Translate(-1.0, 0)
	}
}

func (shape Shape) MoveRight() {
	for _, block := range shape {
		block.GeoM.Translate(1.0, 0)
	}
}

// using '[]' as blocks
func (shape Shape) Draw(screen *ebiten.Image) {
	//draw each coordinate of the tetromino
	new_font := basicfont.Face7x13
	for _, block := range shape {
		text.Draw(screen, "[]", new_font, block.Point_x, block.Point_y, color.RGBA{0xff, 0x00, 0x00, 0x00})
	}
}
