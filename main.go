package main

import (
	"ShitpostBird/dev"
	"ShitpostBird/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth int32 	= 500
	ScreenHeight int32 	= 800
	Title string				= "Shitty Bird"
	FPS int32 					= 60
)

func MainMenu() {
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyEnter) {
			scenes.GameScene()
			break
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		dev.CenterText("Press enter to start", 32, rl.LightGray)
		rl.EndDrawing()
	}
}

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, Title)
	defer rl.CloseWindow()
	
	rl.SetTargetFPS(FPS)

	MainMenu()
	
	if dev.IsGameOver {
		scenes.GameOverScene()
	}
}
