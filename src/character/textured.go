package character

import rl "github.com/gen2brain/raylib-go/raylib"

func (p *Character) GetTexture() *rl.Texture2D {
	return &p.texture
}

func (p *Character) GetTextureDimensions() rl.Vector2 {
	return rl.Vector2{
		X: float32(p.spriteWidth),
		Y: float32(p.spriteHeight),
	}
}
