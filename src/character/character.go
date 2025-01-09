package character

import (
	"anicrossing/src/animation"
	character_animation "anicrossing/src/animation/character"
	"anicrossing/src/inputs"
	"anicrossing/src/state"
	character_state "anicrossing/src/state/character"

	"github.com/gen2brain/raylib-go/raylib"
)

type Character struct {
	texture      rl.Texture2D
	spriteWidth  int
	spriteHeight int

	camera          *rl.Camera2D
	animationPlayer *animation.AnimationPlayer

	position      rl.Vector2
	movementSpeed float32

	state.StateMachine
}

func NewCharacter(
	posX, posY float32,
	camera *rl.Camera2D,
) *Character {
	player := &Character{
		texture:      rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png"),
		spriteWidth:  48,
		spriteHeight: 48,

		camera: camera,

		position:      rl.Vector2{X: posX, Y: posY},
		movementSpeed: 3,
	}

	player.animationPlayer = character_animation.NewCharacterAnimationPlayer(player)

	idleState := character_state.NewCharacterIdleState(player, rl.NewVector2(0, 1))
	walkingState := character_state.NewCharacterWalkingState(player, rl.NewVector2(0, 0))

	player.SetStates(map[string]state.State{
		character_state.Idle:    idleState,
		character_state.Walking: walkingState,
	})

	player.EnterState(idleState)

	return player
}

func (p *Character) Render(delta float32) {
	p.animationPlayer.Render(delta, p.position)
}

func (p *Character) Update() {
	p.camera.Target = rl.NewVector2(
		float32(p.position.X),
		float32(p.position.Y),
	)

	if p.GetCurrentState() == nil {
		return
	}

	p.GetCurrentState().Update()
}

func (p *Character) HandleInput(inputs *inputs.Inputs) {
	if p.GetCurrentState() == nil {
		return
	}

	nextState := p.GetCurrentState().HandleInput(inputs)
	if nextState != nil {
		p.EnterState(nextState)
	}
}

func (p *Character) Move(direction rl.Vector2) {
	p.position.X += direction.X * p.movementSpeed
	p.position.Y += direction.Y * p.movementSpeed
}

func (p *Character) Teardown() {
	rl.UnloadTexture(p.texture)
}
