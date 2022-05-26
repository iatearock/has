// adopted from Hajime Hoshi Ebiten example Blocks
// See https://github.com/hajimehoshi/ebiten/blob/main/examples/blocks/blocks/scenemanager.go

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current            Scene
	next               Scene
	transitionCount    int // frame counter
	transitionMaxCount int // number of frame to transit to next scenen
	transitionFrom     *ebiten.Image
	transitionTo       *ebiten.Image
}

// Create new scene manager, where w - screen width, h - screen height,
// transitionFrames - number of frames to transit between scenes
func NewSceneManager(w, h int) *SceneManager {
	sm := &SceneManager{}
	sm.transitionFrom = ebiten.NewImage(w, h)
	sm.transitionTo = ebiten.NewImage(w, h)
	sm.transitionMaxCount = 20
	return sm
}

func (s *SceneManager) Update() error {
	if s.transitionCount == 0 {
		return s.current.Update()
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.current = s.next
	s.next = nil
	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(screen)
		return
	}

	s.transitionFrom.Clear()
	s.current.Draw(s.transitionFrom)

	s.transitionTo.Clear()
	s.next.Draw(s.transitionTo)

	screen.DrawImage(s.transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(s.transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(s.transitionTo, op)
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = s.transitionMaxCount
	}
}

func (s *SceneManager) SetTransitionFrames(n int) {
	s.transitionMaxCount = n
}
