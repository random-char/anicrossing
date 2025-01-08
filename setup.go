package main

import (
	"anicrossing/src/character"
	"anicrossing/src/tiles"

	"github.com/gen2brain/raylib-go/raylib"
)

var grassTileMap *tiles.TileMap

func setup() {
	rl.InitWindow(screenWidth, screenHeight, "Anicrossing")

	rl.SetTargetFPS(60)

	camera = rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		0.0,
		1.0,
	)
	player = character.NewPlayer(
		screenWidth/2,
		screenHeight/2,
		&camera,
	)

	grassTileMap = tiles.LoadGrassTileMap()

	rl.InitAudioDevice()
	bgMusic = rl.LoadMusicStream("assets/Sound/Music/AveryFarm.mp3")
	rl.PlayMusicStream(bgMusic)
}
