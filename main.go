package main

import (
	"anicrossing/src/character"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

var (
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSpritesheet rl.Texture2D

	bgMusic rl.Music

	player *character.Player
	camera rl.Camera2D
)

func main() {
	setup()

	for running {
		input()
		update()
		render()
	}

	quit()
}
