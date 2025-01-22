package main

import (
	"anicrossing/src/character"
	"anicrossing/src/config"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	conf := config.LoadConfig()

	rl.InitWindow(
		int32(conf.ScreenResolution.X),
		int32(conf.ScreenResolution.Y),
		"map editor",
	)
	rl.SetTargetFPS(60)

	screenCenter := rl.NewVector2((conf.ScreenResolution.X / 2), (conf.ScreenResolution.Y / 2))
	player := character.New(screenCenter)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		player.Render(rl.GetFrameTime())

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
