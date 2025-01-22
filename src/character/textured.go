package character

import rl "github.com/gen2brain/raylib-go/raylib"

func (c *Character) GetTexture() *rl.Texture2D {
	return &c.texture
}

func (c *Character) GetTextureDimensions() rl.Vector2 {
	return rl.Vector2{
		X: float32(c.spriteWidth),
		Y: float32(c.spriteHeight),
	}
}
