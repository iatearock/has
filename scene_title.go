package main

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SceneTitle struct {
	sceneManager *SceneManager
}

func (s *SceneTitle) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return errors.New("0")
	}
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if len(keys) > 0 {
		s.sceneManager.GoTo(NewSceneGame(s.sceneManager))
	}
	return nil
}

func (s *SceneTitle) Draw(screen *ebiten.Image) {
	text.Draw(screen, "Press Any Key to Start", font24, 50, screenHeight/2, color.White)
	text.Draw(screen, "Press Q to quit.", font24, 50, screenHeight/2+30, color.White)
}
