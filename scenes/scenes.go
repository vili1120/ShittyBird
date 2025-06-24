package scenes

import (
	"ShitpostBird/dev"
	"ShitpostBird/scripts"
	etc "ShitpostBird/scripts/Etc"
	pipe "ShitpostBird/scripts/Pipe"
	player "ShitpostBird/scripts/Player"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)


var (
	ScreenWidth int32
	ScreenHeight int32
	Pimg rl.Texture2D
)

func load() {
	ScreenWidth	= int32(rl.GetScreenWidth())
	ScreenHeight = int32(rl.GetScreenHeight())
	Pimg = rl.LoadTexture("assets/images/bird.png")
}

func unload() {
	rl.UnloadTexture(Pimg)
}

func GameScene() {
	load()

	g := &scripts.Game{
		Player: &player.Player{
			Sprite: &etc.Sprite{
				Image: Pimg,
				BB: etc.NewBoundingBox(100, 50, 32, 32),
				V2d: rl.Vector2{X: 100, Y: 50},
			},
			JumpVel: 4.2,
		},
		Pipe1: &pipe.Pipe{
			BB: &pipe.PipeBB{
				TopBB: etc.NewBoundingBox(100, 200, 32, 64),
				BotBB: etc.NewBoundingBox(100, 220, 32, 64),
			},
		},
	}

	g.Pipe1.Target = g.Player

	for !rl.WindowShouldClose() {
		dev.DevUpdate()
		g.Update()
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{130, 163, 255, 100})
		g.Draw()
		if dev.IsGameOver == true {
			break
		}
		rl.EndDrawing()
	}

	unload()
}

func GameOverScene() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		dev.CenterText("Game over!", 32, rl.Red)
		rl.EndDrawing()
	}
}
