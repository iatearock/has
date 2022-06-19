package main

import "github.com/jakecoffman/cp"

type ControlEvent struct {
	exec     map[string]func()
	commands chan string
}

func NewDefualtControlEvent(b *Boat) *ControlEvent {
	r := &ControlEvent{
		exec: map[string]func(){
			"up":     commandUp(b),
			"down":   commandDown(b),
			"left":   commandLeft(b),
			"right":  commandRight(b),
			"anchor": commandAnchor(b),
		},
		commands: make(chan string),
	}
	go r.processEvents()
	return r
}

func (r *ControlEvent) processEvents() {
	for {
		message := <-r.commands
		if message == "up" {
			r.exec["up"]()
		}
		if message == "down" {
			r.exec["down"]()
		}
		if message == "left" {
			r.exec["left"]()
		}
		if message == "right" {
			r.exec["right"]()
		}
		if message == "anchor" {
			r.exec["anchor"]()
		}
	}
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
func commandAnchor(b *Boat) func() {
	return func() {
		if b.isAnchor {
			b.UpAnchor()
		} else {
			b.Anchor()
		}
	}
}
