package player

import (
	"anicrossing/src/character"
	"anicrossing/src/inputs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	camera    *rl.Camera2D
	character *character.Character
}

func New(
	camera *rl.Camera2D,
	character *character.Character,

) *Player {
	return &Player{
		camera:    camera,
		character: character,
	}
}

func (p *Player) Update() {
	p.camera.Target = p.character.GetPosition()

    p.character.Update()
}

func (p *Player) HandleInput(inputs *inputs.Inputs) {
	p.character.HandleInput(inputs)
}

func (p *Player) Teardown() {
	p.character.Teardown()
}
