package main

import (
	"anicrossing/src/character"
	"anicrossing/src/inputs"
	"anicrossing/src/tiles"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	player *character.Character
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

var playerInputs = &inputs.Inputs{
	PressedUp:    false,
	PressedDown:  false,
	PressedLeft:  false,
	PressedRight: false,
	PressedFire:  false,
}

func input() {
	playerInputs.PressedUp = rl.IsKeyDown(rl.KeyW)
	playerInputs.PressedDown = rl.IsKeyDown(rl.KeyS)
	playerInputs.PressedLeft = rl.IsKeyDown(rl.KeyA)
	playerInputs.PressedRight = rl.IsKeyDown(rl.KeyD)
	playerInputs.PressedFire = rl.IsKeyDown(rl.KeySpace)

	player.HandleInput(playerInputs)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)
	rl.BeginMode2D(camera)

	// rl.DrawTexture(grassSpritesheet, 100, 100, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	grassTileMap.Draw(
		tiles.TileGrassTopLeft,
		0,
		0,
	)
	grassTileMap.Draw(
		tiles.TileGrassTopCenter,
		16,
		0,
	)
	grassTileMap.Draw(
		tiles.TileGrassTopCenter,
		32,
		0,
	)

	var tile int

	for row := 0; row < 20; row++ {
		for col := 0; col < 20; col++ {
			if row == 0 {
				if col == 0 {
					tile = tiles.TileGrassTopLeft
				} else if col == 19 {
					tile = tiles.TileGrassTopRight
				} else {
					tile = tiles.TileGrassTopCenter
				}
			} else if row == 19 {
				if col == 0 {
					tile = tiles.TileGrassBottomLeft
				} else if col == 19 {
					tile = tiles.TileGrassBottomRight
				} else {
					tile = tiles.TileGrassBottomCenter
				}
			} else {
				if col == 0 {
					tile = tiles.TileGrassCenterLeft
				} else if col == 19 {
					tile = tiles.TileGrassCenterRight
				} else {
					tile = tiles.TileGrassCenterCenter
				}
			}

			grassTileMap.Draw(
				tile,
				float32(col*16),
				float32(row*16),
			)
		}
	}

	player.Render(rl.GetFrameTime())

	rl.EndMode2D()
	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()

	rl.UpdateMusicStream(bgMusic)

    player.Update()
}

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
	player = character.NewCharacter(
		screenWidth/2,
		screenHeight/2,
		&camera,
	)

	grassTileMap = tiles.LoadGrassTileMap()

	rl.InitAudioDevice()
	bgMusic = rl.LoadMusicStream("assets/Sound/Music/AveryFarm.mp3")
	rl.PlayMusicStream(bgMusic)
}

func quit() {
	player.Teardown()

	rl.UnloadTexture(grassSpritesheet)

	rl.UnloadMusicStream(bgMusic)
	rl.CloseAudioDevice()

	rl.CloseWindow()
}
