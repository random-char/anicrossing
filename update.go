package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update() {
	running = !rl.WindowShouldClose()

	rl.UpdateMusicStream(bgMusic)

    player.Update()
}
