package ebitiles

type Loader struct {
	// TileMap represents the TileMap from which chunks are provided.
	TileMap *TileMap
	// ChunkPos represents the position at specified level.
	// e.g. GlobalPos / (ChunkSize * TileSize)
	ChunkPosX int
	ChunkPosY int
	Level     ChunkLevel
	Radius    int
}

func (l *Loader) Update(x, y int) {
}
