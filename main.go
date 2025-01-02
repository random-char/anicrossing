package main

import (
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

var (
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	tint = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc  rl.Rectangle
	playerDest rl.Rectangle

	playerMovementVector         = rl.Vector2{}
	playerSpeed          float32 = 3
)

func input() {
	playerMovementVector.X = 0
	playerMovementVector.Y = 0

	if rl.IsKeyDown(rl.KeyW) {
		playerMovementVector.Y--
	}
	if rl.IsKeyDown(rl.KeyS) {
		playerMovementVector.Y++
	}
	if rl.IsKeyDown(rl.KeyA) {
		playerMovementVector.X--
	}
	if rl.IsKeyDown(rl.KeyD) {
		playerMovementVector.X++
	}

	playerMovementVector = rl.Vector2Normalize(playerMovementVector)
	playerDest.X += playerMovementVector.X * playerSpeed
	playerDest.Y += playerMovementVector.Y * playerSpeed
}

func update() {
	running = !rl.WindowShouldClose()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)

	rl.DrawTexture(grassSprite, 100, 100, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	rl.DrawTexturePro(
		playerSprite,
		playerSrc,
		playerDest,
		rl.Vector2{X: playerDest.Width, Y: playerDest.Height},
		0,
		tint,
	)

	rl.EndDrawing()
}

func setup() {
	rl.InitWindow(800, 450, "Anicrossing")

	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("assets/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)

	rl.CloseWindow()
}

func main() {
	setup()

	for running {
		input()
		update()
		render()
	}

	quit()
}
