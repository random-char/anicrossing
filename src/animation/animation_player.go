package animation

import (
	"errors"

	"github.com/gen2brain/raylib-go/raylib"
)

type Animated interface {
	GetAnimationPlayer() *AnimationPlayer
}

type AnimationPlayer struct {
	CurrentAnimation *Animation
	Animations       map[string]*Animation
}

func NewAnimationPlayer(
	CurrentAnimation *Animation,
	Animations map[string]*Animation,
) *AnimationPlayer {
	return &AnimationPlayer{
		CurrentAnimation: CurrentAnimation,
		Animations:       Animations,
	}
}

func (ap *AnimationPlayer) Render(delta float32, position rl.Vector2) {
	currAnimation := ap.CurrentAnimation

	currAnimation.cycle(delta)

	rl.DrawTextureRec(
		*currAnimation.Frames[currAnimation.CurrentFrame].Texture,
		currAnimation.Frames[currAnimation.CurrentFrame].Frame,
		position,
		rl.White,
	)
}

func (ap *AnimationPlayer) SetAnimation(name string) error {
	animation, ok := ap.Animations[name]
	if !ok {
		return errors.New("animation not found")
	}

	ap.CurrentAnimation = animation

	return nil
}
