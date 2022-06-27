package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/font"
)

// Message that are located in the world coordinate, i.e. move with other world fixtures
type Message struct {
	pos     cp.Vector // world position
	msg     string
	op      *ebiten.DrawImageOptions
	display bool
}

func NewMessage(text string, pos cp.Vector) *Message {
	m := &Message{pos: pos, msg: text, display: true}
	m.op = &ebiten.DrawImageOptions{}
	m.op.GeoM.Translate(pos.X, pos.Y)
	return m
}

func (m *Message) UpdateText(text string) {
	m.msg = text
}

func (m *Message) UpdatePosition(pos cp.Vector) {
	m.pos = pos
}

func (m *Message) SetDisplay(b bool) {
	m.display = b
}

func (m *Message) Draw(screen *ebiten.Image, fontface font.Face, geom ebiten.GeoM) {
	nop := ebiten.GeoM{}
	nop.Concat(geom)
	nop.Concat(m.op.GeoM)
	x, y := nop.Apply(0, 0)
	text.Draw(screen, m.msg, fontface, int(x), int(y), color.White)
}
