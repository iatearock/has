package ui

import "github.com/hajimehoshi/ebiten/v2"

type Anchor struct {
	img *ebiten.Image

	display bool
	sx      float64
	sy      float64

	op *ebiten.DrawImageOptions
}

func NewAnchor(img *ebiten.Image, sx, sy float64) *Anchor {
	a := &Anchor{img: img, sx: sx, sy: sy, display: false}
	a.op = &ebiten.DrawImageOptions{}
	ix, iy := img.Size()
	a.op.GeoM.Translate(sx-float64(ix)/2, sy-float64(iy)/2)
	return a
}

func (a *Anchor) Draw(screen *ebiten.Image) {
	if a.display {
		screen.DrawImage(a.img, a.op)
	}
}

func (a *Anchor) SetDisplay(b bool) {
	a.display = b
}
