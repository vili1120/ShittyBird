package scripts

import (
	pipe "ShitpostBird/scripts/Pipe"
	player "ShitpostBird/scripts/Player"
)

type Game struct {
	Player *player.Player
	Pipe1 *pipe.Pipe
}

func (g *Game) Update() {
	g.Player.Update()
	g.Pipe1.Update()
}

func (g *Game) Draw() {
	g.Player.Draw()
	g.Pipe1.Draw()
}
