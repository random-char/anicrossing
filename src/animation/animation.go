package animation

type Animation struct {
	Period       float32
	TimeElapsed  float32
	CurrentFrame int
	Frames       []AnimationFrame
}

func (a *Animation) Update(delta float32) {
	a.TimeElapsed += delta

	if a.TimeElapsed >= a.Period {
		a.TimeElapsed -= a.Period
	}
}

func (a *Animation) Cycle(delta float32) {
	a.TimeElapsed += delta
	if a.TimeElapsed >= a.Period {
		//instead of %
		for a.TimeElapsed >= a.Period {
			a.TimeElapsed -= a.Period
		}
		//next frame
		a.CurrentFrame = (a.CurrentFrame + 1) % len(a.Frames)
	}
}
