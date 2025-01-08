package state

import "anicrossing/src/inputs"

type State interface {
	Enter()
	HandleInput(*inputs.Inputs) State
	Update()
	Exit()
}
