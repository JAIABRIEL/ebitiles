package tilemap

type Layer3D struct {
	Image   *EbitImage
	GlobalX int
	GlobalY int
}

func (l *Layer3D) Draw(img *EbitImage) {
	img.DrawAt(l.Image, l.GlobalX, l.GlobalY)
}
