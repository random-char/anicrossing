package main

import (
	"anicrossing/src/tiles"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)
	rl.BeginMode2D(camera)

	// rl.DrawTexture(grassSpritesheet, 100, 100, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	grassTileMap.Draw(
		tiles.TileGrassTopLeft,
		0,
		0,
	)
	grassTileMap.Draw(
		tiles.TileGrassTopCenter,
		16,
		0,
	)
	grassTileMap.Draw(
		tiles.TileGrassTopCenter,
		32,
		0,
	)

	var tile int

	for row := 0; row < 20; row++ {
		for col := 0; col < 20; col++ {
			if row == 0 {
				if col == 0 {
					tile = tiles.TileGrassTopLeft
				} else if col == 19 {
					tile = tiles.TileGrassTopRight
				} else {
					tile = tiles.TileGrassTopCenter
				}
			} else if row == 19 {
				if col == 0 {
					tile = tiles.TileGrassBottomLeft
				} else if col == 19 {
					tile = tiles.TileGrassBottomRight
				} else {
					tile = tiles.TileGrassBottomCenter
				}
			} else {
				if col == 0 {
					tile = tiles.TileGrassCenterLeft
				} else if col == 19 {
					tile = tiles.TileGrassCenterRight
				} else {
					tile = tiles.TileGrassCenterCenter
				}
			}

			grassTileMap.Draw(
				tile,
				float32(col*16),
				float32(row*16),
			)
		}
	}

	player.Render(rl.GetFrameTime())

	rl.EndMode2D()
	rl.EndDrawing()
}
