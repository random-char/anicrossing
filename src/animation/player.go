package animation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	//todo move to player?
	PlayerSpriteWidth  = 48
	PlayerSpriteHeight = 48

	PlayerIdleUpAnimation       = "player-idle-up-animation"
	PlayerIdleDownAnimation     = "player-idle-down-animation"
	PlayerIdleLeftAnimation     = "player-idle-left-animation"
	PlayerIdleRightAnimation    = "player-idle-right-animation"
	PlayerWalkingUpAnimation    = "player-walking-up-animation"
	PlayerWalkingDownAnimation  = "player-walking-down-animation"
	PlayerWalkingLeftAnimation  = "player-walking-left-animation"
	PlayerWalkingRightAnimation = "player-walking-right-animation"
)

type Textured interface {
	GetTexture() *rl.Texture2D
}

func NewPlayerAnimationPlayer(player Textured) *AnimationPlayer {
	idleUpAnimation := createIdleAnimation(player.GetTexture(), 0, -1)
	idleDownAnimation := createIdleAnimation(player.GetTexture(), 0, 1)
	idleLeftAnimation := createIdleAnimation(player.GetTexture(), -1, 0)
	idleRightAnimation := createIdleAnimation(player.GetTexture(), 1, 0)

	return &AnimationPlayer{
		CurrentAnimation: idleDownAnimation,
		Animations: map[string]*Animation{
			PlayerIdleUpAnimation:    idleUpAnimation,
			PlayerIdleDownAnimation:  idleDownAnimation,
			PlayerIdleLeftAnimation:  idleLeftAnimation,
			PlayerIdleRightAnimation: idleRightAnimation,

			PlayerWalkingUpAnimation: createWalkingAnimation(
				player.GetTexture(),
				0, -1,
			),
			PlayerWalkingDownAnimation: createWalkingAnimation(
				player.GetTexture(),
				0, 1,
			),
			PlayerWalkingLeftAnimation: createWalkingAnimation(
				player.GetTexture(),
				-1, 0,
			),
			PlayerWalkingRightAnimation: createWalkingAnimation(
				player.GetTexture(),
				1, 0,
			),
		}}
}

func createIdleAnimation(
	texture *rl.Texture2D,
	facingX, facingY int,
) *Animation {
	var y float32
	if facingY != 0 {
		if facingY > 0 {
			y = 0
		} else {
			y = 1
		}
	} else {
		if facingX < 0 {
			y = 2
		} else {
			y = 3
		}
	}

	frames := make([]AnimationFrame, 2)
	for i := 0; i < 2; i++ {
		frames[i] = AnimationFrame{
			Texture: texture,
			Frame: rl.Rectangle{
				X:      float32(PlayerSpriteWidth * i),
				Y:      float32(PlayerSpriteHeight * y),
				Width:  PlayerSpriteWidth,
				Height: PlayerSpriteHeight,
			},
		}
	}

	return &Animation{
		Period:       0.25,
		TimeElapsed:  0.0,
		CurrentFrame: 0,
		Frames:       frames,
	}
}

func createWalkingAnimation(
	texture *rl.Texture2D,
	facingX, facingY int,
) *Animation {
	var y float32
	if facingY != 0 {
		if facingY > 0 {
			y = 0
		} else {
			y = 1
		}
	} else {
		if facingX < 0 {
			y = 2
		} else {
			y = 3
		}
	}

	frames := make([]AnimationFrame, 4)
	for i := 0; i < 4; i++ {
		frames[i] = AnimationFrame{
			Texture: texture,
			Frame: rl.Rectangle{
				X:      float32(PlayerSpriteWidth * i),
				Y:      float32(PlayerSpriteHeight * y),
				Width:  PlayerSpriteWidth,
				Height: PlayerSpriteHeight,
			},
		}
	}

	return &Animation{
		Period:       0.25,
		TimeElapsed:  0.0,
		CurrentFrame: 0,
		Frames:       frames,
	}
}
