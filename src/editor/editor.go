package editor

import "anicrossing/src/tiles"

type LevelEditor struct {
	DisplayGrid       bool
	AvailableTileMaps map[string]*tiles.TileMap
}

func New() *LevelEditor {
	tileMaps := map[string]*tiles.TileMap{
		"grass": tiles.LoadGrassTileMap(),
	}

	return &LevelEditor{
		DisplayGrid:       true,
		AvailableTileMaps: tileMaps,
	}
}

func (le *LevelEditor) Update() {

}

func (le *LevelEditor) Render() {

}
