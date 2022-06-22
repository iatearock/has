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
	img          *ebiten.Image
}

func NewSceneTitle(sm *SceneManager) *SceneTitle {
	st := &SceneTitle{sceneManager: sm}
	img, _ := vfs.GetImage("assets/images/titlepage.png")
	st.img = img
	return st
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
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.img, op)
	text.Draw(screen, "Run Aground", font24, 50, screenHeight/2-30, color.Black)
	text.Draw(screen, "Press Any Key to Start", font24, 50, screenHeight/2, color.Black)
	text.Draw(screen, "Press Q to quit.", font24, 50, screenHeight/2+30, color.Black)
}
