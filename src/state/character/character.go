package state

import (
	"anicrossing/src/animation"
	"anicrossing/src/movement"
	"anicrossing/src/state"
)

const (
	Idle    = "state-idle"
	Walking = "state-walking"
)

type StatefulAndAnimatedCharacter interface {
	state.Stateful
	animation.Animated
}

type StatefulAndAnimatedMovingCharacter interface {
	state.Stateful
	animation.Animated
    movement.Moving
}
