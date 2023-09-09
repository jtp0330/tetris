package block

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Block struct {
	point [2]int
	img   *ebiten.Image
	geoM  *ebiten.GeoM
}

type Shape [4]*Block

func CreateBlock(img *ebiten.Image, geoM *ebiten.GeoM, tx int, ty int) *Block {
	//create block image
	blockImage := Block{[2]int{tx, ty}, img, geoM}
	return &blockImage
}

// create block randomly in 4 different ways:
func Rotate(shape Shape) (rotated_shape Shape) {
	//block shift left with an XOR operation
	//block shift right with an OR operation
	//block shift down with an AND operation
	//block shift up with an AND NOT operation
	for _, block := range shape {
		block.point[0], block.point[1] = block.point[0]^1, block.point[1]^1
	}

	return
}

func Quick_Down(grid [][]string, block [4]Block) {
	//block shift left with an XOR operation
	//block shift down with an AND operation
	fmt.Println("TBD")
}

func (shape Shape) Draw(screen *ebiten.Image) {
	//draw each coordinate of the tetromino
	new_font := basicfont.Face7x13
	for _, block := range shape {
		text.Draw(screen, "[]", new_font, block.point[0], block.point[1], color.RGBA{0xff, 0x00, 0x00, 0x00})
	}
}

func (shape Shape) MoveDown(y float64) {
	for _, block := range shape {
		block.geoM.Translate(0, y)
	}
}
