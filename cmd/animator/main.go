package main

import (
	"anicrossing/src/character"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(1000, 500, "map editor")
	rl.SetTargetFPS(60)

	camera := rl.NewCamera2D(
		rl.NewVector2(500, 250),
		rl.NewVector2(500, 250),
		0.0,
		1.0,
	)
	player := character.NewCharacter(100, 100, &camera)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		player.Render(rl.GetFrameTime())

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
