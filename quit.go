package main

import rl "github.com/gen2brain/raylib-go/raylib"

func quit() {
	player.Teardown()

	rl.UnloadTexture(grassSpritesheet)

	rl.UnloadMusicStream(bgMusic)
	rl.CloseAudioDevice()

	rl.CloseWindow()
}
