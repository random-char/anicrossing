package character

import (
	"anicrossing/src/animation"
	"anicrossing/src/inputs"
	"anicrossing/src/state"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	idle    = "idle"
	walking = "walking"
)

var (
	playerSpritesheet rl.Texture2D

	spriteWidth  int = 48
	spriteHeight int = 48
	col              = 0
	row              = 0
)

type Player struct {
	texture         rl.Texture2D
	camera          *rl.Camera2D
	animationPlayer *animation.AnimationPlayer

	currSpriteCol int
	currSpriteRow int

	position       rl.Vector2
	movementVector rl.Vector2
	movementSpeed  float32

	currentState state.State
	states       map[string]state.State
}

func NewPlayer(
	posX, posY float32,
	camera *rl.Camera2D,
) *Player {
	player := &Player{
		texture: rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png"),
		camera:  camera,

		position:       rl.Vector2{X: posX, Y: posY},
		movementVector: rl.Vector2{X: 0, Y: 0},
		movementSpeed:  3,
	}

	player.animationPlayer = animation.NewPlayerAnimationPlayer(player)

	idleState := &PlayerIdleState{
		player: player,
	}
	walkingState := &PlayerWalkingState{
		player: player,
		dx:     0,
		dy:     0,
	}

	player.states = map[string]state.State{
		idle:    idleState,
		walking: walkingState,
	}

	player.enterState(idleState)

	return player
}

func (p *Player) GetTexture() *rl.Texture2D {
	return &p.texture
}

func (p *Player) Render(delta float32) {
	p.animationPlayer.Render(delta, p.position)
}

func (p *Player) Update() {
	p.camera.Target = rl.NewVector2(
		float32(p.position.X),
		float32(p.position.Y),
	)

	if p.currentState == nil {
		return
	}

	p.currentState.Update()
}

func (p *Player) HandleInput(inputs *inputs.Inputs) {
	if p.currentState == nil {
		return
	}

	nextState := p.currentState.HandleInput(inputs)
	if nextState != nil {
		p.enterState(nextState)
	}
}

func (p *Player) enterState(state state.State) {
	if p.currentState != nil {
		p.currentState.Exit()
	}

	p.currentState = state
	p.currentState.Enter()
}

func (p *Player) Teardown() {
	rl.UnloadTexture(playerSpritesheet)
}
