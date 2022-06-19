package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputHandler struct {
	keys       map[string]ebiten.Key // change this to change control
	eventQueue chan string           // pass input to eventQueue
}

// connect to event processing channel
func NewInputHandler(eq chan string) *InputHandler {
	return &InputHandler{keys: map[string]ebiten.Key{},
		eventQueue: eq}
}

func NewDefaultInputHandler(eq chan string) *InputHandler {
	ih := &InputHandler{keys: map[string]ebiten.Key{
		"up":     ebiten.KeyUp,
		"down":   ebiten.KeyDown,
		"left":   ebiten.KeyLeft,
		"right":  ebiten.KeyRight,
		"anchor": ebiten.KeyD,
	},
		eventQueue: eq,
	}

	return ih
}

func (ih *InputHandler) Update() error {
	if ebiten.IsKeyPressed(ih.keys["up"]) {
		ih.eventQueue <- "up"
	}
	if ebiten.IsKeyPressed(ih.keys["down"]) {
		ih.eventQueue <- "down"
	}
	if ebiten.IsKeyPressed(ih.keys["left"]) {
		ih.eventQueue <- "left"
	}
	if ebiten.IsKeyPressed(ih.keys["right"]) {
		ih.eventQueue <- "right"
	}
	if inpututil.IsKeyJustPressed(ih.keys["anchor"]) {
		ih.eventQueue <- "anchor"
	}
	return nil
}
