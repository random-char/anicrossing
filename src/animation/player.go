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
	idleAnimation := &Animation{
		Period:       0.5,
		TimeElapsed:  0.0,
		CurrentFrame: 0,
		Frames: []AnimationFrame{
			{
				Texture: player.GetTexture(),
				Frame: rl.Rectangle{
					X:      0,
					Y:      0,
					Width:  PlayerSpriteWidth,
					Height: PlayerSpriteHeight,
				},
			},
			{
				Texture: player.GetTexture(),
				Frame: rl.Rectangle{
					X:      PlayerSpriteWidth,
					Y:      0,
					Width:  PlayerSpriteWidth,
					Height: PlayerSpriteHeight,
				},
			},
		},
	}
	walkingAnimation := &Animation{
		Period:       0.5,
		TimeElapsed:  0.0,
		CurrentFrame: 0,
		Frames: []AnimationFrame{
			{
				Texture: player.GetTexture(),
				Frame: rl.Rectangle{
					X:      0,
					Y:      PlayerSpriteHeight * 2,
					Width:  PlayerSpriteWidth,
					Height: PlayerSpriteHeight,
				},
			},
			{
				Texture: player.GetTexture(),
				Frame: rl.Rectangle{
					X:      PlayerSpriteWidth,
					Y:      PlayerSpriteHeight * 2,
					Width:  PlayerSpriteWidth,
					Height: PlayerSpriteHeight,
				},
			},
			{
				Texture: player.GetTexture(),
				Frame: rl.Rectangle{
					X:      PlayerSpriteWidth * 2,
					Y:      PlayerSpriteHeight * 2,
					Width:  PlayerSpriteWidth,
					Height: PlayerSpriteHeight,
				},
			},
			{
				Texture: player.GetTexture(),
				Frame: rl.Rectangle{
					X:      PlayerSpriteWidth * 3,
					Y:      PlayerSpriteHeight * 2,
					Width:  PlayerSpriteWidth,
					Height: PlayerSpriteHeight,
				},
			},
		},
	}

	return &AnimationPlayer{
		CurrentAnimation: idleAnimation,
		Animations: map[string]*Animation{
			PlayerIdleUpAnimation:    idleAnimation,
			PlayerIdleDownAnimation:  idleAnimation,
			PlayerIdleLeftAnimation:  idleAnimation,
			PlayerIdleRightAnimation: idleAnimation,

			PlayerWalkingUpAnimation:    walkingAnimation,
			PlayerWalkingDownAnimation:  walkingAnimation,
			PlayerWalkingLeftAnimation:  walkingAnimation,
			PlayerWalkingRightAnimation: walkingAnimation,
		}}
}
