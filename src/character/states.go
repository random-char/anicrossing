package character

import (
	"anicrossing/src/animation"
	"anicrossing/src/inputs"
	"anicrossing/src/state"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerIdleState struct {
	player  *Player
	facingX int
	facingY int
}

func (idle *PlayerIdleState) Enter() {
	fmt.Println("enter idle")
}

func (idle *PlayerIdleState) Exit() {
	fmt.Println("exit idle")
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
	player *Player
	dx     float32
	dy     float32
}

func (walking *PlayerWalkingState) Enter() {
	var walkingAnimation string

	if walking.dy != 0 {
		if walking.dy > 0 {
			walkingAnimation = animation.PlayerWalkingUpAnimation
		} else {
			walkingAnimation = animation.PlayerWalkingDownAnimation
		}
	} else {
		if walking.dx > 0 {
			walkingAnimation = animation.PlayerWalkingRightAnimation
		} else {
			walkingAnimation = animation.PlayerWalkingLeftAnimation
		}
	}

	walking.player.animationPlayer.SetAnimation(walkingAnimation)
}

func (walking *PlayerWalkingState) Exit() {
	fmt.Println("exit walking")
}

func (walking *PlayerWalkingState) HandleInput(inputs *inputs.Inputs) state.State {
	if !inputs.DirectionPressed() {
		idleState, ok := walking.player.states[idle].(*PlayerIdleState)
		if !ok {
			panic("idle state set wrongly")
		}
		idleState.facingX = int(walking.dx)
		idleState.facingY = int(walking.dy)

		return idleState
	}

	walking.dx = 0
	walking.dy = 0

	if inputs.PressedRight {
		walking.dx++
	}
	if inputs.PressedLeft {
		walking.dx--
	}
	if inputs.PressedUp {
		walking.dy--
	}
	if inputs.PressedDown {
		walking.dy++
	}

	return nil
}

func (walking *PlayerWalkingState) Update() {
	walking.player.movementVector.X = walking.dx
	walking.player.movementVector.Y = walking.dy

	walking.player.movementVector = rl.Vector2Normalize(walking.player.movementVector)

	walking.player.position.X += walking.player.movementVector.X * walking.player.movementSpeed
	walking.player.position.Y += walking.player.movementVector.Y * walking.player.movementSpeed
}
