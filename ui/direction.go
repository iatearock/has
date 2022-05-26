package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

// Direction show a pointer at the edge of screen indicates where to go
type Direction struct {
	img     *ebiten.Image
	target  cp.Vector
	current cp.Vector

	display bool // do not display when target is on screen
	sx      float64
	sy      float64
	op      *ebiten.DrawImageOptions
}

// Create a new direction image `img`, to show on screen at
// position `sx`, `sy`, pointing to `target`
func NewDirection(img *ebiten.Image, sx, sy float64,
	target cp.Vector) *Direction {
	return &Direction{img: img, sx: sx, sy: sy,
		target: target, current: cp.Vector{}, display: true,
		op: &ebiten.DrawImageOptions{},
	}
}

func (d *Direction) SetPosition(v cp.Vector) {
	d.current = v
	d.op.GeoM.Reset()
	ix, iy := d.img.Size()
	d.op.GeoM.Translate(float64(-ix/2), float64(-iy/2))
	dir := d.target.Sub(d.current)
	angle := dir.ToAngle()
	d.op.GeoM.Rotate(angle)
	d.op.GeoM.Translate(d.sx, d.sy)
}

func (d *Direction) SetTarget(v cp.Vector) {
	d.target = v
}

func (d *Direction) Draw(screen *ebiten.Image) {
	if d.display {
		screen.DrawImage(d.img, d.op)
	}
}
