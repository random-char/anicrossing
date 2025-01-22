package character

import "anicrossing/src/animation"

func (c *Character) GetAnimationPlayer() *animation.AnimationPlayer {
	return c.animationPlayer
}
