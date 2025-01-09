package player

import (
	"anicrossing/src/character"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	camera    *rl.Camera2D
	character *character.Character
}
