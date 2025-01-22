package state

import (
	character_animation "anicrossing/src/animation/character"
	"anicrossing/src/inputs"
	"anicrossing/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CharacterWalkingState struct {
	character StatefulAndAnimatedMovingCharacter
	direction rl.Vector2
}

func NewCharacterWalkingState(
	character StatefulAndAnimatedMovingCharacter,
	direction rl.Vector2,
) *CharacterWalkingState {
	return &CharacterWalkingState{
		character: character,
		direction: direction,
	}
}

func (walking *CharacterWalkingState) Enter() {}

func (walking *CharacterWalkingState) Exit() {
	walking.direction.X = 0
	walking.direction.Y = 0
}

func (walking *CharacterWalkingState) HandleInput(inputs *inputs.Inputs) state.State {
	var X, Y float32 = 0, 0
	if inputs.PressedRight {
		X++
	}
	if inputs.PressedLeft {
		X--
	}
	if inputs.PressedUp {
		Y--
	}
	if inputs.PressedDown {
		Y++
	}

	if X == 0 && Y == 0 {
		state := walking.character.GetState(Idle)
		switch idleState := state.(type) {
		case *CharacterIdleState:
			//set idle facing to last movement direction
			idleState.facing.X = walking.direction.X
			idleState.facing.Y = walking.direction.Y
		}

		return state
	}

	if X != walking.direction.X || Y != walking.direction.Y {
		//update animation
		var walkingAnimation string

		if Y != 0 {
			if Y > 0 {
				walkingAnimation = character_animation.CharacterWalkingDownAnimation
			} else {
				walkingAnimation = character_animation.CharacterWalkingUpAnimation
			}
		} else {
			if X > 0 {
				walkingAnimation = character_animation.CharacterWalkingRightAnimation
			} else {
				walkingAnimation = character_animation.CharacterWalkingLeftAnimation
			}
		}

		walking.character.GetAnimationPlayer().SetAnimation(walkingAnimation)
	}

	walking.direction.X = X
	walking.direction.Y = Y

	return nil
}

func (walking *CharacterWalkingState) Update() {
	walking.character.Move(rl.Vector2Normalize(walking.direction))
}
