package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	pos         Position
	targetPos   Position
	path        []Position
	moving      bool
	moveCounter int
	score       int
	sprite      *ebiten.Image
	direction   int // 1 для правого движения, -1 для левого
}

func NewPlayer(sprite *ebiten.Image) Player {
	return Player{
		pos:       Position{0, 0},
		sprite:    sprite,
		direction: 1,
	}
}

func (p *Player) UpdatePosition(gridWidth, gridHeight int) {
	if p.moving && len(p.path) > 0 {
		p.moveCounter++
		if p.moveCounter >= moveSpeed {
			nextPos := p.path[0]
			// Определяем направление движения
			if nextPos.X > p.pos.X {
				p.direction = 1 // Движется вправо
			} else if nextPos.X < p.pos.X {
				p.direction = -1 // Движется влево
			}

			p.pos = nextPos
			p.path = p.path[1:]
			p.moveCounter = 0

			if len(p.path) == 0 && p.pos != p.targetPos {
				p.path = FindPath(p.pos, p.targetPos, gridWidth, gridHeight)
				p.moving = len(p.path) > 0
			} else if len(p.path) == 0 {
				p.moving = false
			}
		}
	}
}

func (p *Player) SetTarget(targetPos Position, gridWidth, gridHeight int) {
	if targetPos != p.targetPos {
		p.targetPos = targetPos
		p.path = FindPath(p.pos, p.targetPos, gridWidth, gridHeight)
		p.moving = len(p.path) > 0
		p.moveCounter = 0
	}
}

func (p *Player) CollectPrizes(prizes []Prize) []Prize {
	for i := 0; i < len(prizes); {
		if prizes[i].pos == p.pos {
			p.score++
			prizes = append(prizes[:i], prizes[i+1:]...)
		} else {
			i++
		}
	}
	return prizes
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Масштабируем изображение игрока до размера ячейки
	frameWidth := p.sprite.Bounds().Dx()
	frameHeight := p.sprite.Bounds().Dy()

	scaleX := float64(cellSize) / float64(frameWidth)
	scaleY := float64(cellSize) / float64(frameHeight)

	op.GeoM.Scale(scaleX, scaleY)

	// Поворот в зависимости от направления
	if p.direction == -1 {
		// Разворачиваем на 180 градусов по оси Y
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(cellSize), 0) // Сдвигаем обратно в видимую область
	}

	op.GeoM.Translate(float64(p.pos.X*cellSize), float64(p.pos.Y*cellSize))

	screen.DrawImage(p.sprite, op)
}
