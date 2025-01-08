package inputs

type Inputs struct {
	PressedUp    bool
	PressedDown  bool
	PressedLeft  bool
	PressedRight bool
	PressedFire  bool
}

func (i *Inputs) DirectionPressed() bool {
	return i.PressedUp ||
		i.PressedDown ||
		i.PressedLeft ||
		i.PressedRight
}
