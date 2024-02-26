package tilemap

type Layer3D struct {
	Level int
	Image TileImage
	X     int
	Y     int
}

func (l *Layer3D) Draw(img TileImage) {
	img.DrawAt(img, l.X, l.Y)
}
