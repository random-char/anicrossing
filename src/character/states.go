package character

import (
	"anicrossing/src/animation"
	"anicrossing/src/inputs"
	"anicrossing/src/state"

	"github.com/gen2brain/raylib-go/raylib"
)

type PlayerIdleState struct {
	player *Player
	facing rl.Vector2
}

func (idle *PlayerIdleState) Enter() {
	var idleAnimation string

	if idle.facing.Y != 0 {
		if idle.facing.Y > 0 {
			idleAnimation = animation.PlayerIdleDownAnimation
		} else {
			idleAnimation = animation.PlayerIdleUpAnimation
		}
	} else {
		if idle.facing.X < 0 {
			idleAnimation = animation.PlayerIdleLeftAnimation
		} else {
			idleAnimation = animation.PlayerIdleRightAnimation
		}
	}

	idle.player.animationPlayer.SetAnimation(idleAnimation)
}

func (idle *PlayerIdleState) Exit() {
	idle.facing.X = 0
	idle.facing.Y = 0
}

func (idle *PlayerIdleState) HandleInput(inputs *inputs.Inputs) state.State {
	if inputs.DirectionPressed() {
		return idle.player.states[walking]
	}

	return nil
}

func (idle *PlayerIdleState) Update() {
}

type PlayerWalkingState struct {
	player    *Player
	direction rl.Vector2
}

func (walking *PlayerWalkingState) Enter() {
}

func (walking *PlayerWalkingState) Exit() {
	walking.direction.X = 0
	walking.direction.Y = 0
}

func (walking *PlayerWalkingState) HandleInput(inputs *inputs.Inputs) state.State {
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
		idleState, ok := walking.player.states[idle].(*PlayerIdleState)
		if !ok {
			panic("idle state set wrongly")
		}
		//set idle facing to last movement direction
		idleState.facing.X = walking.direction.X
		idleState.facing.Y = walking.direction.Y

		return idleState
	}

	if X != walking.direction.X || Y != walking.direction.Y {
		//update animation
		var walkingAnimation string

		if Y != 0 {
			if Y > 0 {
				walkingAnimation = animation.PlayerWalkingDownAnimation
			} else {
				walkingAnimation = animation.PlayerWalkingUpAnimation
			}
		} else {
			if X > 0 {
				walkingAnimation = animation.PlayerWalkingRightAnimation
			} else {
				walkingAnimation = animation.PlayerWalkingLeftAnimation
			}
		}

		walking.player.animationPlayer.SetAnimation(walkingAnimation)
	}

	walking.direction.X = X
	walking.direction.Y = Y

	return nil
}

func (walking *PlayerWalkingState) Update() {
	walking.direction = rl.Vector2Normalize(walking.direction)

	walking.player.position.X += walking.direction.X * walking.player.movementSpeed
	walking.player.position.Y += walking.direction.Y * walking.player.movementSpeed
}
