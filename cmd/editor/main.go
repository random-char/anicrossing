package main

import (
	"anicrossing/src/editor"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1000, 500, "map editor")
    rl.SetTargetFPS(60)

    levelEditor := editor.New()

	for !rl.WindowShouldClose() {
        rl.BeginDrawing()

        levelEditor.Update()
        levelEditor.Render()

        rl.EndDrawing()
	}

	rl.CloseWindow()
}
