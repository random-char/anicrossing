package main

import (
	"anicrossing/src/character"
	"anicrossing/src/inputs"
	"anicrossing/src/player"
	"anicrossing/src/tiles"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

var (
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSpritesheet rl.Texture2D
	grassTileMap     *tiles.TileMap

	bgMusic rl.Music

	characterInstance *character.Character
	camera            rl.Camera2D
	playerInstance    *player.Player
	playerInputs      *inputs.Inputs = &inputs.Inputs{
		PressedUp:    false,
		PressedDown:  false,
		PressedLeft:  false,
		PressedRight: false,
		PressedFire:  false,
	}
)

func main() {
	setup()

	for !rl.WindowShouldClose() {
		input()
		update()
		render()
	}

	quit()
}

func input() {
	playerInputs.PressedUp = rl.IsKeyDown(rl.KeyW)
	playerInputs.PressedDown = rl.IsKeyDown(rl.KeyS)
	playerInputs.PressedLeft = rl.IsKeyDown(rl.KeyA)
	playerInputs.PressedRight = rl.IsKeyDown(rl.KeyD)
	playerInputs.PressedFire = rl.IsKeyDown(rl.KeySpace)

	characterInstance.HandleInput(playerInputs)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)
	rl.BeginMode2D(camera)

	renderBackgroundStub()
	characterInstance.Render(rl.GetFrameTime())

	rl.EndMode2D()
	rl.EndDrawing()
}

func update() {
	rl.UpdateMusicStream(bgMusic)

	characterInstance.Update()
}

func setup() {
	rl.InitWindow(screenWidth, screenHeight, "Anicrossing")

	rl.SetTargetFPS(60)

	camera = rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		0.0,
		1.0,
	)
	characterInstance = character.NewCharacter(
		screenWidth/2,
		screenHeight/2,
		&camera,
	)
	playerInstance = player.NewPlayer(&camera, characterInstance)

	grassTileMap = tiles.LoadGrassTileMap()

	rl.InitAudioDevice()
	bgMusic = rl.LoadMusicStream("assets/Sound/Music/AveryFarm.mp3")

	rl.PlayMusicStream(bgMusic)
}

func quit() {
	characterInstance.Teardown()

	rl.UnloadTexture(grassSpritesheet)
	rl.UnloadMusicStream(bgMusic)

	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func renderBackgroundStub() {
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
}
