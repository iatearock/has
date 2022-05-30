package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type InputHandler struct {
	keys     map[string]ebiten.Key // change this to change control
	commands map[ebiten.Key]func()
}

func NewInputHandler() *InputHandler {
	return &InputHandler{keys: map[string]ebiten.Key{}}
}

func NewDefaultInputHandler(b *Boat) *InputHandler {
	ih := &InputHandler{keys: map[string]ebiten.Key{
		"up":    ebiten.KeyUp,
		"down":  ebiten.KeyDown,
		"left":  ebiten.KeyLeft,
		"right": ebiten.KeyRight,
	},
	}
	ih.commands = map[ebiten.Key]func(){
		ih.keys["up"]:    commandUp(b),
		ih.keys["down"]:  commandDown(b),
		ih.keys["left"]:  commandLeft(b),
		ih.keys["right"]: commandRight(b),
	}

	return ih
}

func (ih *InputHandler) Update() error {
	if ebiten.IsKeyPressed(ih.keys["up"]) {
		ih.commands[ih.keys["up"]]()
	}
	if ebiten.IsKeyPressed(ih.keys["down"]) {
		ih.commands[ih.keys["down"]]()
	}
	if ebiten.IsKeyPressed(ih.keys["left"]) {
		ih.commands[ih.keys["left"]]()
	}
	if ebiten.IsKeyPressed(ih.keys["right"]) {
		ih.commands[ih.keys["right"]]()
	}
	return nil
}

func commandUp(b *Boat) func() {
	return func() {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 100, Y: 0}, cp.Vector{})
	}
}
func commandDown(b *Boat) func() {
	return func() {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: -100, Y: 0}, cp.Vector{})
	}
}
func commandLeft(b *Boat) func() {
	return func() {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: 50}, cp.Vector{X: -20, Y: 0})
	}
}
func commandRight(b *Boat) func() {
	return func() {
		b.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -50}, cp.Vector{X: -20, Y: 0})
	}
}
