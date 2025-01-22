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

	animationPlayer *animation.AnimationPlayer

	position      rl.Vector2
	movementSpeed float32

	state.StateMachine
}

func New(
	position rl.Vector2,
) *Character {
	character := &Character{
		texture:      rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png"),
		spriteWidth:  48,
		spriteHeight: 48,

		position: position,
        movementSpeed: 3,
	}

	character.animationPlayer = character_animation.NewCharacterAnimationPlayer(character)

	idleState := character_state.NewCharacterIdleState(character, rl.NewVector2(0, 1))
	walkingState := character_state.NewCharacterWalkingState(character, rl.NewVector2(0, 0))

	character.SetStates(map[string]state.State{
		character_state.Idle:    idleState,
		character_state.Walking: walkingState,
	})

	character.EnterState(idleState)

	return character
}

func (c *Character) GetPosition() rl.Vector2 {
	return c.position
}

func (c *Character) Render(delta float32) {
	c.animationPlayer.Render(delta, c.position)
}

func (c *Character) Update() {
	if c.GetCurrentState() == nil {
		return
	}

	c.GetCurrentState().Update()
}

func (c *Character) HandleInput(inputs *inputs.Inputs) {
	if c.GetCurrentState() == nil {
		return
	}

	nextState := c.GetCurrentState().HandleInput(inputs)
	if nextState != nil {
		c.EnterState(nextState)
	}
}

func (c *Character) Move(direction rl.Vector2) {
	c.position.X += direction.X * c.movementSpeed
	c.position.Y += direction.Y * c.movementSpeed
}

func (c *Character) Teardown() {
	rl.UnloadTexture(c.texture)
}
