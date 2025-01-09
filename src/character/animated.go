package character

import "anicrossing/src/animation"

func (p *Character) GetAnimationPlayer() *animation.AnimationPlayer {
	return p.animationPlayer
}
