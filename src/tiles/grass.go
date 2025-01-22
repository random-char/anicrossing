package tiles

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TileGrassTopLeft      = 0
	TileGrassTopCenter    = 1
	TileGrassTopRight     = 2
	TileGrassCenterLeft   = 11
	TileGrassCenterCenter = 12
	TileGrassCenterRight  = 13
	TileGrassBottomLeft   = 22
	TileGrassBottomCenter = 23
	TileGrassBottomRight  = 24
)

var TileNames map[int]string = map[int]string{
	TileGrassTopLeft: "grass-top-left",
}

func LoadGrassTileMap() *TileMap {
	tileMap := &TileMap{
		Texture:    rl.LoadTexture("assets/Tilesets/Grass.png"),
		TileWidth:  16,
		TileHeight: 16,
		Columns:    11,
		Rows:       7,
	}

	tiles := make([]*Tile, 77)

	i := 0
	for c := 0; c < tileMap.Columns; c++ {
		for r := 0; r < tileMap.Rows; r++ {
			if (c == 10 && (r < 2 || r > 3)) ||
				(c == 9 && r == 6) {
				i++
				//skip empty
				continue
			}

			tiles[i] = &Tile{
				X:     float32(c),
				Y:     float32(r),
				Solid: false,
			}
			i++
		}
	}

	tileMap.Tiles = tiles

	return tileMap
}
