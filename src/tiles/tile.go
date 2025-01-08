package tiles

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileMap struct {
	Texture    rl.Texture2D
	TileWidth  float32
	TileHeight float32
	Columns    int
	Rows       int
	Tiles      []*Tile
}

type Tile struct {
	X     float32
	Y     float32
	Solid bool
}

func (t *TileMap) Draw(
	tile int,
	x, y float32,
) {
	tileRow := tile / int(t.Columns)
	tileCol := tile % int(t.Columns)

	rl.DrawTexturePro(
		t.Texture,
		rl.NewRectangle(
			float32(tileCol)*t.TileWidth,
			float32(tileRow)*t.TileHeight,
			t.TileWidth,
			t.TileHeight,
		),
		rl.NewRectangle(x, y, t.TileWidth, t.TileHeight),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}
