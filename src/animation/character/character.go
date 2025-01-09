package character

import (
	"anicrossing/src/animation"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CharacterIdleUpAnimation       = "player-idle-up-animation"
	CharacterIdleDownAnimation     = "player-idle-down-animation"
	CharacterIdleLeftAnimation     = "player-idle-left-animation"
	CharacterIdleRightAnimation    = "player-idle-right-animation"
	CharacterWalkingUpAnimation    = "player-walking-up-animation"
	CharacterWalkingDownAnimation  = "player-walking-down-animation"
	CharacterWalkingLeftAnimation  = "player-walking-left-animation"
	CharacterWalkingRightAnimation = "player-walking-right-animation"
)

type Textured interface {
	GetTexture() *rl.Texture2D
	GetTextureDimensions() rl.Vector2
}

func NewCharacterAnimationPlayer(player Textured) *animation.AnimationPlayer {
	vectorUp := rl.NewVector2(0, -1)
	vectorDown := rl.NewVector2(0, 1)
	vectorLeft := rl.NewVector2(-1, 0)
	vectorRight := rl.NewVector2(1, 0)

	animations := map[string]*animation.Animation{
		CharacterIdleUpAnimation:    createIdleAnimation(player, vectorUp),
		CharacterIdleDownAnimation:  createIdleAnimation(player, vectorDown),
		CharacterIdleLeftAnimation:  createIdleAnimation(player, vectorLeft),
		CharacterIdleRightAnimation: createIdleAnimation(player, vectorRight),

		CharacterWalkingUpAnimation:    createWalkingAnimation(player, vectorUp),
		CharacterWalkingDownAnimation:  createWalkingAnimation(player, vectorDown),
		CharacterWalkingLeftAnimation:  createWalkingAnimation(player, vectorLeft),
		CharacterWalkingRightAnimation: createWalkingAnimation(player, vectorRight),
	}

	return animation.NewAnimationPlayer(
		animations[CharacterIdleDownAnimation],
		animations,
	)
}

func createIdleAnimation(
	character Textured,
	facing rl.Vector2,
) *animation.Animation {
	var y float32 = 0
	if facing.Y != 0 {
		if facing.Y < 0 {
			y = 1
		}
	} else {
		if facing.X < 0 {
			y = 2
		} else {
			y = 3
		}
	}

	spriteWidth := character.GetTextureDimensions().X
	spriteHeight := character.GetTextureDimensions().Y

	frames := make([]animation.AnimationFrame, 2)
	for i := 0; i < 2; i++ {
		frames[i] = animation.AnimationFrame{
			Texture: character.GetTexture(),
			Frame: rl.Rectangle{
				X:      spriteWidth * float32(i),
				Y:      spriteHeight * float32(y),
				Width:  spriteWidth,
				Height: spriteHeight,
			},
		}
	}

	return &animation.Animation{
		Period:       0.25,
		TimeElapsed:  0.0,
		CurrentFrame: 0,
		Frames:       frames,
	}
}

func createWalkingAnimation(
	character Textured,
	facing rl.Vector2,
) *animation.Animation {
	var y float32
	if facing.Y != 0 {
		if facing.Y < 0 {
			y = 1
		}
	} else {
		if facing.X < 0 {
			y = 2
		} else {
			y = 3
		}
	}

	spriteWidth := character.GetTextureDimensions().X
	spriteHeight := character.GetTextureDimensions().Y

	frames := make([]animation.AnimationFrame, 4)
	for i := 0; i < 4; i++ {
		frames[i] = animation.AnimationFrame{
			Texture: character.GetTexture(),
			Frame: rl.Rectangle{
				X:      spriteWidth * float32(i),
				Y:      spriteHeight * float32(y),
				Width:  spriteWidth,
				Height: spriteHeight,
			},
		}
	}

	return &animation.Animation{
		Period:       0.25,
		TimeElapsed:  0.0,
		CurrentFrame: 0,
		Frames:       frames,
	}
}
