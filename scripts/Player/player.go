package player

import (
	"ShitpostBird/dev"
	etc "ShitpostBird/scripts/Etc"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	*etc.Sprite
	JumpVel float64
	Yvel float64
}

func (p *Player) Update() {
	p.V2d.Y += float32(p.Yvel)
	p.Yvel += 0.1

	if rl.IsKeyPressed(rl.KeySpace) {
		p.Yvel = -p.JumpVel
	}
	if p.IsTouchingEdge() {
		dev.IsGameOver = true
	}
}

func (p *Player) Draw() {
	p.BB.Rect.X = p.V2d.X+(float32(p.Image.Width)/2)-(p.BB.Rect.Width/2)
	p.BB.Rect.Y = p.V2d.Y+(float32(p.Image.Height)/2)-(p.BB.Rect.Height/2)
	rl.DrawTexture(p.Image, int32(p.V2d.X), int32(p.V2d.Y), rl.White)
	if dev.Debug == 1 {
		rl.DrawRectangleLinesEx(p.BB.Rect, 1, rl.Green)
	}
}

func (p *Player) IsTouchingEdge() bool {
	return p.V2d.Y <= 0 || p.Image.Height+int32(p.V2d.Y) >= int32(rl.GetScreenHeight())
}
