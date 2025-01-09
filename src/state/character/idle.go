package state

import (
	"anicrossing/src/animation"
	character_animation "anicrossing/src/animation/character"
	"anicrossing/src/inputs"
	"anicrossing/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StatefulAndAnimatedCharacter interface {
	state.Stateful
	animation.Animated
}

type CharacterIdleState struct {
	character StatefulAndAnimatedCharacter
	facing    rl.Vector2
}

func NewCharacterIdleState(
	character StatefulAndAnimatedCharacter,
	facing rl.Vector2,
) *CharacterIdleState {
	return &CharacterIdleState{
		character: character,
		facing:    facing,
	}
}

func (idle *CharacterIdleState) Enter() {
	var idleAnimation string

	if idle.facing.Y != 0 {
		if idle.facing.Y > 0 {
			idleAnimation = character_animation.CharacterIdleDownAnimation
		} else {
			idleAnimation = character_animation.CharacterIdleUpAnimation
		}
	} else {
		if idle.facing.X < 0 {
			idleAnimation = character_animation.CharacterIdleLeftAnimation
		} else {
			idleAnimation = character_animation.CharacterIdleRightAnimation
		}
	}

	idle.character.GetAnimationPlayer().SetAnimation(idleAnimation)
}

func (idle *CharacterIdleState) HandleInput(inputs *inputs.Inputs) state.State {
	if inputs.DirectionPressed() {
		return idle.character.GetStates()[Walking]
	}

	return nil
}

func (idle *CharacterIdleState) Update() {}

func (idle *CharacterIdleState) Exit() {
	idle.facing.X = 0
	idle.facing.Y = 0
}
