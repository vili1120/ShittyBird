package dev

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Debug = 1

var IsGameOver = false

func DevUpdate() {
	if rl.IsKeyPressed(rl.KeyRightShift) {
		Debug++
	}
	if Debug == 2 {
		Debug = 0
	}
}

func CenterText(message string, fontSize int, color rl.Color) {
	msgWidth := rl.MeasureText(message, int32(fontSize))
	cx := (rl.GetScreenWidth()-int(msgWidth))/2
	cy := (rl.GetScreenHeight()-fontSize)/2
	rl.DrawText(message, int32(cx), int32(cy), int32(fontSize), color)
}
