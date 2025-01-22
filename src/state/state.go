package state

import "anicrossing/src/inputs"

type Stateful interface {
	GetStates() map[string]State
	GetState(string) State
	GetCurrentState() State
	EnterState(State)
}

type State interface {
	Enter()
	HandleInput(*inputs.Inputs) State
	Update()
	Exit()
}

type StateMachine struct {
	currentState State
	states       map[string]State
}

func (sm *StateMachine) SetStates(states map[string]State) {
	sm.states = states
}

func (sm *StateMachine) GetStates() map[string]State {
	return sm.states
}

func (sm *StateMachine) GetState(name string) State {
	return sm.states[name]
}

func (sm *StateMachine) GetCurrentState() State {
	return sm.currentState
}

func (sm *StateMachine) EnterState(state State) {
	if sm.currentState != nil {
		sm.currentState.Exit()
	}

	sm.currentState = state
	sm.currentState.Enter()
}
