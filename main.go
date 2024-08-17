package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1024
	screenHeight = 720
	cellSize     = 75
	moveSpeed    = 5
)

type Game struct {
	player       Player
	prizeManager PrizeManager
	gridWidth    int
	gridHeight   int
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		newTargetPos := Position{mouseX / cellSize, mouseY / cellSize}
		g.player.SetTarget(newTargetPos, g.gridWidth, g.gridHeight)
	}

	g.player.UpdatePosition(g.gridWidth, g.gridHeight)

	// Проверка на сбор приза
	if g.prizeManager.CollectPrize(g.player.pos) {
		// Если приз собран, он будет удален, а новый приз появится автоматически
		g.player.score++
	}
	g.prizeManager.Update(g.gridWidth, g.gridHeight)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Отрисовка layout
	DrawGrid(screen, screenWidth, screenHeight, cellSize)
	// Отрисовка игрока и цели
	g.player.Draw(screen)

	g.prizeManager.Draw(screen)

	g.drawScore(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) drawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Score: %d", g.player.score)
	ebitenutil.DebugPrintAt(screen, scoreText, screenWidth-150, 20)
}

func main() {
	// Загрузка изображений для призов
	prizeSprites := []*ebiten.Image{}
	for _, file := range []string{"sneaker0.png", "sneaker2.png", "sneaker4.png"} {
		img, _, err := ebitenutil.NewImageFromFile(file)
		if err != nil {
			log.Fatalf("Error loading prize image: %v", err)
		}
		prizeSprites = append(prizeSprites, img)
	}

	// Загрузка спрайта игрока
	sprite, _, err := ebitenutil.NewImageFromFile("roma8.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		player:       NewPlayer(sprite),
		gridWidth:    screenWidth / cellSize,
		gridHeight:   screenHeight / cellSize,
		prizeManager: NewPrizeManager(prizeSprites),
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("2D Grid Movement Example")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
