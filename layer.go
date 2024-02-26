package tilemap

type Layer3D struct {
	Image TileImage
	X     int
	Y     int
}

func (l *Layer3D) Draw(img TileImage) {
	img.DrawAt(l.Image, l.X, l.Y)
}
