package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	prizeLifetime = 10 // Время жизни приза (в секундах)
)

type Prize struct {
	pos       Position
	spawnTime time.Time
	sprite    *ebiten.Image
}

type PrizeManager struct {
	prize   *Prize // У нас всегда есть один активный приз
	sprites []*ebiten.Image
}

func (pm *PrizeManager) Update(gridWidth, gridHeight int) {
	now := time.Now()

	if pm.prize == nil {
		// Если приза нет, создаем его
		pm.spawnNewPrize(gridWidth, gridHeight)
	} else if now.Sub(pm.prize.spawnTime).Seconds() > prizeLifetime {
		// Если прошло 10 секунд, приз появляется в новом месте
		pm.spawnNewPrize(gridWidth, gridHeight)
	}
}

func (pm *PrizeManager) spawnNewPrize(gridWidth, gridHeight int) {
	spriteIndex := rand.Intn(len(pm.sprites)) // Выбираем случайное изображение
	pm.prize = &Prize{
		pos:       Position{X: rand.Intn(gridWidth), Y: rand.Intn(gridHeight)},
		spawnTime: time.Now(),
		sprite:    pm.sprites[spriteIndex],
	}
}

func (pm *PrizeManager) CollectPrize(playerPos Position) bool {
	if pm.prize != nil && pm.prize.pos == playerPos {
		pm.prize = nil // Приз собран, удаляем его
		return true
	}
	return false
}

func (pm *PrizeManager) Draw(screen *ebiten.Image) {
	if pm.prize != nil {
		op := &ebiten.DrawImageOptions{}

		// Масштабируем изображение приза до размера ячейки
		frameWidth := pm.prize.sprite.Bounds().Dx()
		frameHeight := pm.prize.sprite.Bounds().Dy()

		scaleX := float64(cellSize) / float64(frameWidth)
		scaleY := float64(cellSize) / float64(frameHeight)

		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(pm.prize.pos.X*cellSize), float64(pm.prize.pos.Y*cellSize))

		screen.DrawImage(pm.prize.sprite, op)
	}
}

func NewPrizeManager(sprites []*ebiten.Image) PrizeManager {
	return PrizeManager{
		sprites: sprites,
	}
}
