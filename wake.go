package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type Bubble struct {
	pos      cp.Vector
	size     float64 // default image is 100px dia circle
	grow     float64 // grow rate in diameter per sec
	lifetime float64 // life limit in sec
	age      float64 // life in sec
	alpha    float64 // transparency
	fade     float64 // fade rate for alph
}

// NewBubble at pos, with initial size and boat velocity
func NewBubble(pos cp.Vector, size, vel float64) *Bubble {
	lt := 5.0
	initAlpha := 127.0
	return &Bubble{pos: pos, size: size, grow: vel * 0.5,
		lifetime: lt, age: 0, alpha: initAlpha, fade: initAlpha / lt}
}

func (b *Bubble) Update() {
	dt := 1 / 60.0
	b.age += dt
	b.size += b.grow * dt
	b.alpha -= b.fade * dt
}

func (b *Bubble) IsAlive() bool {
	return b.age < b.lifetime
}

func (b *Bubble) Draw(screen, bb *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-50, -50) // HARDCODE image size
	sx := 1.0 / 100.0 * b.size
	op.GeoM.Scale(sx, sx)
	op.GeoM.Translate(b.pos.X, b.pos.Y)
	op.GeoM = game.cam.Concat(op.GeoM)
	op.ColorM.Scale(1, 1, 1, b.alpha/127.0)
	screen.DrawImage(bb, op)
}

type Wake struct {
	img     *ebiten.Image
	bubbles []*Bubble
}

func NewWake() *Wake {
	img, err := vfs.GetImage("assets/images/bubble100.png")
	if err != nil {
		log.Println(err)
	}
	return &Wake{img: img, bubbles: []*Bubble{}}
}

func (w *Wake) AddBubble(pos cp.Vector, size, vel float64) {
	w.bubbles = append(w.bubbles, NewBubble(pos, size, vel))
}
func (w *Wake) Update() {
	for _, b := range w.bubbles {
		b.Update()
	}
	i := 0
	for _, b := range w.bubbles {
		if b.IsAlive() { // remove out of date bubble
			w.bubbles[i] = b
			i++
		}
	}
	w.bubbles = w.bubbles[:i]
}

func (w *Wake) Draw(screen *ebiten.Image) {
	for _, b := range w.bubbles {
		b.Draw(screen, w.img)
		// TODO filter out of screen
	}
}
