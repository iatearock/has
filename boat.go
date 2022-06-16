package main

import (
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type Boat struct {
	hull   *ebiten.Image
	sail   *ebiten.Image // maybe a list of images
	rudder *ebiten.Image

	rudderAngle float64 // radian -ve to turn right
	sailAngle   float64 // radian -ve to turn right
	sailUp      float64 // percentage, 0 to 1

	imageW, imageH   int // boat image width/height
	imageWr, imageHr int // rudder image width/height
	imageWs, imageHs int // sail image width/height

	body  *cp.Body
	shape *cp.Shape

	inStore   bool // is boat in store
	store     []*Item
	maxWeight float64 // boat store weight capacity

	bubbleCount int // counter for emitting wake
}

func NewBoat(x, y float64) *Boat {
	h, errH := vfs.GetImage("assets/images/dinghy.png")
	if errH != nil {
		log.Println(errH)
	}
	r, errR := vfs.GetImage("assets/images/rudder.png")
	if errR != nil {
		log.Println(errR)
	}
	s, errS := vfs.GetImage("assets/images/sail.png")
	if errS != nil {
		log.Println(errS)
	}
	imageW, imageH := h.Size()
	imageWr, imageHr := r.Size()
	imageWs, imageHs := s.Size()
	b := &Boat{hull: h, sail: s, rudder: r,
		imageW: imageW, imageH: imageH,
		imageWr: imageWr, imageHr: imageHr,
		imageWs: imageWs, imageHs: imageHs,
		sailUp: 0.0, sailAngle: 0.0, rudderAngle: 0.0,
		store: []*Item{}, maxWeight: 50, inStore: false,
		bubbleCount: 0,
	}
	b.body = space.AddBody(cp.NewBody(10, cp.MomentForBox(10, 40, 20))) // mass, moment
	b.body.SetPosition(cp.Vector{X: x, Y: y})
	b.shape = space.AddShape(cp.NewCircle(b.body, 10, cp.Vector{}))
	// TODO better body and shape
	b.shape.SetCollisionType(COLLISION_BOAT)
	b.body.UserData = b

	return b
}

func (b *Boat) Draw(screen *ebiten.Image) {
	pos := b.body.Position()
	// hull
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(b.imageW/2), -float64(b.imageH/2))
	op.GeoM.Rotate(b.body.Angle())
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM = game.cam.Concat(op.GeoM)
	screen.DrawImage(b.hull, op)
	// rudder
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(b.imageWr), -float64(b.imageHr/2))
	op.GeoM.Rotate(b.rudderAngle)
	op.GeoM.Translate(-float64(b.imageW/2), 0)
	op.GeoM.Rotate(b.body.Angle())
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM = game.cam.Concat(op.GeoM)
	screen.DrawImage(b.rudder, op)
	// sail
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(b.imageWs), -float64(b.imageHs/2))
	op.GeoM.Rotate(b.sailAngle)
	op.GeoM.Translate(float64(b.imageW/4), 0)
	op.GeoM.Rotate(b.body.Angle())
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM = game.cam.Concat(op.GeoM)
	screen.DrawImage(b.sail, op)

	// text.Draw(screen, fmt.Sprintf("Sail %d%% Up", int(b.sailUp*100)), font24, 10, 40, color.White)
}

func (b *Boat) Update() error {
	// b.ControlRow()
	// b.ControlWind()

	// Lateral resistance
	vecY := cp.ForAngle(b.body.Angle() + math.Pi/2)
	velY := b.body.Velocity().Dot(vecY)
	if velY > 1 {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -math.Pow(velY, 2.0)}, cp.Vector{})
	}
	if velY < -1 {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: math.Pow(velY, 2.0)}, cp.Vector{})
	}

	b.bubbleCount++
	if b.bubbleCount > 60 {
		b.bubbleCount = 0
	}
	if b.bubbleCount%30 == 0 {
		off := rand.Float64()*20.0 - 10.0
		yOffset := vecY.Normalize().Mult(off)
		game.wake.AddBubble(b.body.Position().Add(yOffset), 5, b.body.Velocity().Length())
	}

	return nil
}

func (b *Boat) ControlWind() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if b.sailUp < 1.0 {
			b.sailUp += 0.1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if b.sailUp > 0.0 {
			b.sailUp -= 0.1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if b.rudderAngle > -math.Pi/2 {
			b.rudderAngle -= math.Pi / 8
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if b.rudderAngle < math.Pi/2 {
			b.rudderAngle += math.Pi / 8
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyPeriod) {
		if b.sailAngle > -math.Pi/2 {
			b.sailAngle -= math.Pi / 8
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyComma) {
		if b.sailAngle < math.Pi/2 {
			b.sailAngle += math.Pi / 8
		}
	}
}
func (b *Boat) ControlRow() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 100, Y: 0}, cp.Vector{})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: -100, Y: 0}, cp.Vector{})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -50}, cp.Vector{X: -20, Y: 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: 50}, cp.Vector{X: -20, Y: 0})
	}
}

// RemovePhysics body and shape before deleting the boat
func (b *Boat) RemovePhysics() {
	space.RemoveShape(b.shape)
	space.RemoveBody(b.body)
}

func (b *Boat) String() string {
	return "boat"
}

func (b *Boat) EnterStore() {
	b.inStore = true
}

func (b *Boat) ExitStore() {
	b.inStore = false
}
