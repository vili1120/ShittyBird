package pipe

import (
	"ShitpostBird/dev"
	etc "ShitpostBird/scripts/Etc"
	player "ShitpostBird/scripts/Player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PipeBB struct {
	TopBB *etc.BoundingBox
	BotBB *etc.BoundingBox
}

func (pbb *PipeBB) CheckIfCollidedWithPlayer(p *player.Player) {
	if pbb.TopBB.IsColliding(p.BB.Rect) || pbb.BotBB.IsColliding(p.BB.Rect) {
		dev.IsGameOver = true
	}
}

func (pbb *PipeBB) DrawHitBoxes() {
	if pbb.TopBB != nil {
		rl.DrawRectangleLinesEx(pbb.TopBB.Rect, 1.5, rl.Red)
	}
	if pbb.BotBB != nil {
		rl.DrawRectangleLinesEx(pbb.BotBB.Rect, 1.5, rl.Red)
	}
}

type Pipe struct {
	BB *PipeBB
	Target *player.Player
}

func (p *Pipe) Update() {
	p.BB.CheckIfCollidedWithPlayer(p.Target)
}

func (p *Pipe) Draw() {
	if dev.Debug == 1 {
		p.BB.DrawHitBoxes()
	}
}
