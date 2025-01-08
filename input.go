package main

import (
	"anicrossing/src/inputs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
