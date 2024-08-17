package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// DrawGrid отрисовывает сетку на экране
func DrawGrid(screen *ebiten.Image, screenWidth, screenHeight, cellSize int) {
	screen.Fill(color.RGBA{211, 211, 211, 255}) // light gray

}
