package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type Island struct {
	img     *ebiten.Image
	overlay *ebiten.Image
	radius  float64

	body   *cp.Body
	shape  *cp.Shape // collision
	sensor *cp.Shape // is boat in range to dock

	store []*Item
	// offer []int // offer/ask price for store item
	// ask   []int

	end bool // finish island for this level
}

// Create new circular island at x, y with radius r
func NewIsland(x, y, r float64, end bool) *Island {
	img, err := vfs.GetImage("assets/images/island.png")

	if err != nil {
		log.Println(err)
	}
	imgOverlay, err := vfs.GetImage("assets/images/treasure.png")

	if err != nil {
		log.Println(err)
	}
	i := &Island{img: img, overlay: imgOverlay, radius: r, store: []*Item{}, end: end}
	i.body = space.AddBody(cp.NewStaticBody())
	i.body.SetPosition(cp.Vector{X: x, Y: y})
	i.sensor = space.AddShape(cp.NewCircle(i.body, r+20, cp.Vector{}))
	i.sensor.SetSensor(true)
	i.sensor.SetCollisionType(COLLISION_SENSOR)
	i.shape = space.AddShape(cp.NewCircle(i.body, r, cp.Vector{}))

	i.body.UserData = i
	return i
}

func (i *Island) Draw(screen *ebiten.Image) {
	w, h := i.img.Size()
	sx := i.radius / float64(w) * 2
	sy := i.radius / float64(h) * 2

	pos := i.body.Position()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-i.radius, -i.radius)
	op.GeoM.Scale(sx, sy)
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM = game.cam.Concat(op.GeoM)
	screen.DrawImage(i.img, op)
	if i.end {
		screen.DrawImage(i.overlay, op)
	}
}

// RemovePhysics body and shape before deleting the island
func (i *Island) RemovePhysics() {
	space.RemoveShape(i.shape)
	space.RemoveShape(i.sensor)
	space.RemoveBody(i.body)
}

func (i *Island) String() string {
	return "island"
}
