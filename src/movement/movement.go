package movement

import rl "github.com/gen2brain/raylib-go/raylib"

type Moving interface {
	Move(rl.Vector2)
}
