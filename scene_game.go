package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"iatearock.com/has/ui"
)

type SceneGame struct {
	sceneManager *SceneManager
	direction    *ui.Direction
	inputHandler *InputHandler
}

func NewSceneGame(sm *SceneManager) *SceneGame {
	s := &SceneGame{sceneManager: sm}
	dirImg, _ := vfs.GetImage("assets/images/direction.png")
	s.direction = ui.NewDirection(dirImg, 600, 40, game.end.body.Position())
	// game.audioManager.audioPlayer.Play()
	wave, _ := game.audioManager.NewInfiniteLoop("wave", "assets/audio/632517__thedutchmancreative__waves.ogg")
	seagull, _ := game.audioManager.NewPlayer("seagull", "assets/audio/510917__lydmakeren__seagulls-short.ogg")
	wave.Play()
	seagull.Play()
	s.inputHandler = NewDefaultInputHandler(game.b)
	return s
}

func (s *SceneGame) Update() error {
	s.inputHandler.Update()
	game.b.Update()
	space.Step(1.0 / 60.0)

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		s.sceneManager.GoTo(&SceneTitle{s.sceneManager})
		game.audioManager.CloseAll()
	}
	game.wake.Update()

	// update camera and HUD
	bp := game.b.body.Position()
	s.direction.SetPosition(bp) // HUD arrow pointer
	game.cam.Position = [2]float64{bp.X, bp.Y}
	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{128, 170, 187, 255})

	game.cam.Update()
	text.Draw(screen, "Press q to return to Title Screen", font24, 50, screenHeight/8*7, color.White)
	game.island.Draw(screen)
	game.start.Draw(screen)
	game.end.Draw(screen)
	game.wake.Draw(screen)
	game.b.Draw(screen)

	s.direction.Draw(screen)
	s.DrawInfo(screen)
}

func (s *SceneGame) DrawInfo(screen *ebiten.Image) {
	if game.b.inStore {
		text.Draw(screen, "Shop", font24, 500, 40, color.White)
	}
}
