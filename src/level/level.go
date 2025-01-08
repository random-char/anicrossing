package level

import "anicrossing/src/tiles"


//todo add serialization
type Level struct {
    TileWidth int
    TileHeight int
    TilePlacement []*tiles.Tile
}
